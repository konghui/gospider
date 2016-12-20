package tuicool

import (
	"log"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/konghui/gospider/task"
	"github.com/konghui/gospider/urltool"
)

func sendrequest(url string) (response *http.Response, err error) {
	var mytask task.ClientTask

	mytask.Url, err = urltool.NewURL(url)
	if err != nil {
		return
	}
	response, err = site.SendHttpRequest(&mytask)

	return
}

func Test_ProcessResponse(t *testing.T) {
	//var url = "http://www.tuicool.com/articles/jaER3eZ"
	var url = "http://www.tuicool.com/articles/yUVBRjf"
	response, err := sendrequest(url)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	list, err := ParseUrlToList(doc)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	t.Logf("get the house list:\n")
	for i, v := range list {
		t.Logf("%d:%s\n", i, v)
	}

	web := NewTuiCool()
	err = web.GetContent(doc)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf(web.String())
}
