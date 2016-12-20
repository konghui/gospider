package tuicool

import "github.com/konghui/gospider/web"

var site = web.NewWebSite("www.tuicool.com")

/*
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
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	list, err = ParseUrlToList(doc)
	if err != nil {
		return
	}
	return
}
*/
