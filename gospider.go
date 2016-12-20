package main

import (
	"fmt"
	"regexp"

	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/client"
	"github.com/konghui/gospider/proto"
	"github.com/konghui/gospider/queue"
	"github.com/konghui/gospider/web"
)

var BaseUrl = "http://www.cnblogs.com/tt-0411/archive/2013/03/13/2958130.html"

var content = (`(?s).*<div class="postBody">(.*)<div id="blog_post_info_block">.*`)
var contentPartern = regexp.MustCompile(content)

//

func GetHttpUrlList() (task string) {
	task = queue.MyQueue.Pop()
	return
}

func StartConsumer() {
	client := client.NewClient()
	site := web.NewWebSite("http://bj.58.com/changping/chuzu/0/b12/")
	site.RegisterRule(`(?s).*<div class="postBody">(.*)<div id="blog_post_info_block">.*`)
	client.Site["bj.58.com"] = site
	client.ConsumerLoop()
}

func KeepAliveHandler(args *proto.Request, rv *proto.Response) {
	//rv = proto.NewKeepAliveResponse()
	rv.BuildKeepAliveResponse()
}

func ParsePageHandler(args *proto.Request, rv *proto.Response) {
	rv.BuildParsePageResponse("www.google.com")
}

func ReportDataHandler(args *proto.Request, rv *proto.Response) {
	fmt.Println("report data")
}

func GetUrlListHandler(args *proto.Request, rv *proto.Response) {
	rv.BuildGetListResponse([]string{"www.google.com", "www.baidu.com"})
}

func StartProducer() {
	log.Info("Start Producer")
	server, err := proto.NewServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	server.RegisterCallbackFunc(proto.Ping, KeepAliveHandler)
	server.RegisterCallbackFunc(proto.ParsePage, ParsePageHandler)
	server.RegisterCallbackFunc(proto.ReportData, ReportDataHandler)
	server.RegisterCallbackFunc(proto.GetUrlList, GetUrlListHandler)

	queue.MyQueue.Append(BaseUrl)
	ProducerLoop()
}

func ProducerLoop() {
	for {
		if queue.MyQueue.Length() == 0 {
			break
		}
	}
}

func main() {
	StartProducer()
	go StartConsumer()
	select {}
}
