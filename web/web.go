package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/check"
	"github.com/konghui/gospider/filter"
	"github.com/konghui/gospider/queue"
	"github.com/konghui/gospider/task"
	"github.com/konghui/gospider/urltool"
)

type SiteHandler interface {
	ParseUrlToList(*goquery.Document) ([]string, error)
	GetContent(*goquery.Document) error
}

type WebSite struct {
	domain     string
	filterRule []*filter.FilterRule
	client     *http.Client
	handler    SiteHandler
	check      check.CheckUrl
}

func NewWebSite(domain string) (web *WebSite) {

	web = &WebSite{domain: domain, filterRule: make([]*filter.FilterRule, 0), client: &http.Client{}}
	log.Infof("create new website  %s", domain)
	return
}

func (this *WebSite) RegisterCheckHandler(handler check.CheckUrl) {
	this.check = handler
}

func (this *WebSite) RegisterHandler(handler SiteHandler) {
	this.handler = handler
}

func (this *WebSite) GetFilterRule() (fr []*filter.FilterRule) {
	return this.filterRule
}

func (this *WebSite) RegisterRule(rule string) {

	filterRule := filter.NewFilterRule(rule)
	this.filterRule = append(this.filterRule, filterRule)
	log.Infof("register new rule %s for website %s.", rule, this.domain)
}

func (this *WebSite) SendHttpRequest(task *task.ClientTask) (response *http.Response, err error) {

	var request *http.Request

	log.Infof("Send http request to url:%s, method:%s.", task.Url, http.MethodGet)
	request, err = http.NewRequest(http.MethodGet, task.Url.StandUrl(), nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/53.0.2785.143 Chrome/53.0.2785.143 Safari/537.36")
	if err != nil {
		return
	}
	response, err = this.client.Do(request)

	if err != nil {
		return
	}

	log.Infof("get the response from the server.")
	return
}

func GetResponseContent(response *http.Response) (content string, err error) {
	var out []byte
	defer response.Body.Close()
	out, err = ioutil.ReadAll(response.Body)
	content = string(out)
	if err != nil {
		return
	}
	log.Debugf("response body:%s", content)
	return
}

func (this *WebSite) Work() {
	var e error
	var t task.ClientTask
	for {
		url := queue.MyQueue.Pop()
		if url == "" {
			log.Warning("all the work has done")
			break
		}
		log.Infof("start url: %s", url)
		t.Url, e = urltool.NewURL(url)
		if e != nil {
			log.Warning(e.Error())
		}
		this.check.SetString(t.Url.String())
		response, err := this.SendHttpRequest(&t)
		defer response.Body.Close()
		if err != nil {
			log.Warning(err.Error())
		}
		doc, err := goquery.NewDocumentFromResponse(response)
		if err != nil {
			log.Fatal(err.Error())
		}
		list, err := this.handler.ParseUrlToList(doc)
		if err != nil {
			log.Warningf(err.Error())
			return
		}

		for _, v := range list {
			newUrl := fmt.Sprintf("http://%s/%s", "www.tuicool.com", v)
			//fmt.Printf("%d:%s\n", i, newUrl)
			myurl, err := urltool.NewURL(newUrl)
			if err != nil {
				log.Warning(err.Error())
			}
			if this.check.GetString(myurl.String()) {
				log.Infof("url:%s already visit. skip it", myurl)
				continue
			}
			queue.MyQueue.Append(newUrl)
			this.check.SetString(myurl.String())
		}

		err = this.handler.GetContent(doc)
		if err != nil {
			log.Warningf(err.Error())
		}
		log.Info(this.handler)

		time.Sleep(1 * time.Second)
	}
}
