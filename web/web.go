package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/check"
	"github.com/konghui/gospider/filter"
	"github.com/konghui/gospider/queue"
	"github.com/konghui/gospider/task"
)

type SiteHandler interface {
	String() string
	GetHandler() (fun []func(*SpiderArgs) error)
}

type WebSite struct {
	domain     string
	filterRule []*filter.FilterRule
	client     *http.Client
	handler    SiteHandler
	visitPool  *check.VisitPool
}

func NewWebSite(domain string) (web *WebSite, err error) {

	web = &WebSite{domain: domain, filterRule: make([]*filter.FilterRule, 0), client: &http.Client{}}
	err = web.SetVisitPool("bloomfileter", uint64(1024), "murmur3Hash64", "jenkinsHash64")
	if err != nil {
		return
	}
	log.Infof("create new website  %s", domain)
	return
}

func (this *WebSite) SetVisitPool(name string, args ...interface{}) (err error) {
	this.visitPool, err = check.NewVisitPool(name, args)
	return
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

	url := task.GetUrl()
	log.Infof("Send http request to url:%s, method:%s.", url, http.MethodGet)
	request, err = http.NewRequest(http.MethodGet, url.StandUrl(), nil)
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

// read the text string from the response
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

// convert the related URI of the page (which visit by the task *t) to the URL
func GetURLOfTheTaskPageLink(t *task.ClientTask, link string) (u string) {
	if strings.HasPrefix(link, "/") { // link startwith '/' was the related url, such as /index.html
		url := t.GetUrl()
		u = fmt.Sprintf("%s://%s:%d%s", url.Proto, url.Domain, url.Port, link)
	}

	return
}

func (this *WebSite) NewSpiderArgs(url string) (args *SpiderArgs) {

	log.Infof("start url: %s", url)
	task, err := task.NewTask(url)
	this.visitPool.Visit(url)
	response, err := this.SendHttpRequest(task)
	defer response.Body.Close()
	if err != nil {
		log.Warning(err.Error())
	}

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	args = &SpiderArgs{Doc: doc, Response: response, VisitMap: this.visitPool, Queue: queue.MyQueue, Task: task}
	return
}

func (this *WebSite) StartLoop() {
	for {

		url := queue.MyQueue.Pop()
		if url == "" {
			log.Warning("all the work has done")
			break
		}
		args := this.NewSpiderArgs(url)
		for _, handler := range this.handler.GetHandler() {
			err := handler(args)
			if err != nil {
				log.Warningf(err.Error())
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}
