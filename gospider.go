package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var BaseUrl = "http://www.baidu.com"
var client = &http.Client{}

func sendHttpRequest(url string) (list []string, err error) {

	var request *http.Request
	var response *http.Response
	list = make([]string, 0)
	err = nil

	log.Infof("Send http request to url:%s, method:%s.", http.MethodGet, url)
	request, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	response, err = client.Do(request)

	//	response, err = http.ReadResponse(&buffer, request)
	if err != nil {
		return
	}
	log.Infof("get the response.")
	fmt.Println(response.Body)
	return
}

func main() {
	var urllist []string
	var list []string
	var err error
	urllist = append(urllist, BaseUrl)
	for {
		if len(urllist) == 0 {
			break
		}
		url := urllist[0]
		list, err = sendHttpRequest(url)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("*******")
		urllist = urllist[1:len(urllist)]
		fmt.Println(list)
	}
}
