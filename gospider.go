package main

import (
	"fmt"
	"net/http"

	"regexp"

	"io/ioutil"

	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/urltool"
)

var BaseUrl = "http://www.cnblogs.com/tt-0411/archive/2013/03/13/2958130.html"

var content = (`(?s).*<div class="postBody">(.*)<div id="blog_post_info_block">.*`)
var contentPartern = regexp.MustCompile(content)

type Queue struct {
	queue  []string
	mutex  sync.Mutex
	length uint32
}

var queue Queue

type ClientTask struct {
	Url *urltool.URL
}

func (this *Queue) Append(task string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.queue = append(this.queue, task)
	this.length = this.length + 1
	log.Debugf("add new task:%s to the queue", task)
}

func (this *Queue) Pop() (data string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.length == 0 {
		return ""
	}
	data = this.queue[0]
	this.queue = this.queue[1:this.length]
	this.length = this.length - 1
	log.Debugf("pop the task:%s from queue", data)
	return
}

func (this *Queue) Length() (len uint32) {
	return this.length
}

type Client struct {
	httpClient *http.Client
	site       map[string]*WebSite
}

type FilterRule struct {
	rule     string
	parttern *regexp.Regexp
	data     string
}

func NewFilterRule(rule string) (filterRule *FilterRule) {
	filterRule = &FilterRule{rule: rule, parttern: regexp.MustCompile(rule)}
	return
}

type WebSite struct {
	domain     string
	filterRule []*FilterRule
	client     *http.Client
}

func NewWebSite(domain string) (web *WebSite) {

	web = &WebSite{domain: domain, filterRule: make([]*FilterRule, 0), client: &http.Client{}}
	log.Infof("create new website  %s", domain)
	return
}

func (this *WebSite) RegisterRule(rule string) {

	filterRule := NewFilterRule(rule)
	this.filterRule = append(this.filterRule, filterRule)
	log.Infof("register new rule %s for website %s.", rule, this.domain)
}
func NewClient() (client *Client) {
	client = &Client{
		httpClient: &http.Client{},
		site:       make(map[string]*WebSite),
	}
	return
}

func (this *Client) ConsumerLoop() {
	log.Info("Enter the consumer loop")
	for {
		var task ClientTask
		var err error

		url := GetHttpUrlList()
		if url == "" {
			log.Info("url set is empty!")
			time.Sleep(1 * time.Second)
			continue
		}
		task.Url, err = urltool.NewURL(url)
		if err != nil {
			log.Warning(err.Error())
		}
		if web, yes := this.site[task.Url.Domain]; yes == true {
			content, err := web.sendHttpRequest(&task)
			if err != nil {
				log.Warning(err.Error())
				continue
			}
			for _, filter := range web.filterRule {
				groups := filter.parttern.FindStringSubmatch(string(content))
				fmt.Printf("%q", groups)
				if len(groups) < 2 {
					log.Warn("No match found for the rule:%s, page:%s", filter.rule, task.Url.StandUrl())
					continue
				}

				filter.data = groups[1]
				fmt.Println(filter.data)
			}
		}

	}
}

func StartConsumer() {
	client := NewClient()
	site := NewWebSite("http://bj.58.com/changping/chuzu/0/b12/")
	site.RegisterRule(`(?s).*<div class="postBody">(.*)<div id="blog_post_info_block">.*`)
	client.site["bj.58.com"] = site
	client.ConsumerLoop()
}

func GetHttpUrlList() (task string) {
	task = queue.Pop()
	return
}

func (this *WebSite) sendHttpRequest(task *ClientTask) (content []byte, err error) {

	var request *http.Request
	var response *http.Response

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
	defer response.Body.Close()
	content, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	log.Debugf("response body:%s", string(content))

	return
}

func ValidUrl(task ClientTask) (rv bool) {
	return true
}

func StartProducer() {
	log.Info("Start Producer")
	queue.Append(BaseUrl)

}

func ProducerLoop() {
	for {
		if queue.Length() == 0 {
			break
		}

	}

}

func main() {
	StartProducer()
	go StartConsumer()
	select {}
}
