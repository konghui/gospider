package main

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"

	_ "github.com/konghui/gospider/check/bfchecker"
	"github.com/konghui/gospider/proto"
	"github.com/konghui/gospider/queue"
	"github.com/konghui/gospider/web"
	"github.com/konghui/gospider/web/tuicool"
)

const BIND_ADDR = "127.0.0.1:1234"

func CheckAlive() {
	for {
		log.Infof("SendRequest to the %s", BIND_ADDR)
		response, err := proto.SendKeepAliveRequest(BIND_ADDR)
		log.Debug(response)
		if err != nil {
			log.Warn(err.Error())
		}
		if response.IsPongResponse() {
			log.Infof("Get the keeplive response from %s", BIND_ADDR)
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	log.SetLevel(log.InfoLevel)
	//queue.MyQueue.Append("http://www.tuicool.com/articles/UBbIFnA")
	queue.MyQueue.Append("http://www.tuicool.com/")
	//	go CheckAlive()
	website, err := web.NewWebSite("www.tuicool.com")
	if err != nil {
		log.Fatal(err.Error())
	}
	website.RegisterHandler(tuicool.NewTuiCool())
	go website.StartLoop()
	fmt.Println("all work has done")
	select {}
}
