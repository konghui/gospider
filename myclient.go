package main

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/konghui/bloomfilter"
	"github.com/konghui/gospider/proto"
	"github.com/konghui/gospider/queue"
	"github.com/konghui/gospider/web"
	"github.com/konghui/gospider/web/tuicool"
	"github.com/konghui/hashland/jenkins"
	"github.com/konghui/hashland/murmur3"
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

func murmur3Hash64(data []byte) (hashcode uint64) {
	hashcode = murmur3.Sum64(data)
	return
}

func jenkinsHash64(data []byte) (hashcode uint64) {
	hashcode = jenkins.Hash264(data, 0)
	return
}

func initBloom() (bf *bloomfilter.BloomFilter) {
	bf = bloomfilter.New(1024)
	bf.Register(murmur3Hash64)
	bf.Register(jenkinsHash64)
	return
}

func main() {
	log.SetLevel(log.InfoLevel)
	//queue.MyQueue.Append("http://www.tuicool.com/articles/UBbIFnA")
	queue.MyQueue.Append("http://www.tuicool.com/")
	//	go CheckAlive()
	website := web.NewWebSite("www.tuicool.com")
	website.RegisterHandler(tuicool.NewTuiCool())
	website.RegisterCheckHandler(initBloom())
	website.Work()
	fmt.Println("all work has done")
	select {}
}
