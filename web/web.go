package web

import (
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/filter"
	"github.com/konghui/gospider/task"
)

type WebSite struct {
	domain     string
	filterRule []*filter.FilterRule
	client     *http.Client
}

func NewWebSite(domain string) (web *WebSite) {

	web = &WebSite{domain: domain, filterRule: make([]*filter.FilterRule, 0), client: &http.Client{}}
	log.Infof("create new website  %s", domain)
	return
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

	log.Infof("Send http request to url:%s, method:%s.", http.MethodGet, task.Url)
	request, err = http.NewRequest(http.MethodGet, task.Url.StandUrl(), nil)
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

func (this *WebSite) GetResponseContent(response *http.Response) (content string, err error) {
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
