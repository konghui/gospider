package web58

import (
	"net/http"

	"github.com/konghui/gospider/task"
	"github.com/konghui/gospider/urltool"
	"github.com/konghui/gospider/web"
)

var site = web.NewWebSite("bj.58.com")

//var URL = "http://bj.58.com/tiantongyuan/zufang/0/j2/"

func GetUrlList(url string) (list []string, err error) {
	var mytask task.ClientTask
	var response *http.Response
	mytask.Url, err = urltool.NewURL(url)
	if err != nil {
		return
	}
	response, err = site.SendHttpRequest(&mytask)

	if err != nil {
		return
	}
	list, err = GetHouseList(response)
	if err != nil {
		return
	}
	return
}
