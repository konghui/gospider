package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/proto"
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
	go CheckAlive()
	select {}
}
