package tuicool

import (
	"fmt"
	"strings"
	"time"

	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type TuiCool struct {
	// the title of the article
	title string
	// the tag of the article
	label []string
	// the unix time stamp of the article
	timestamp int64
	// the origin of the article
	origin string
	// the content of the article
	content string
}

var list = make([]string, 0)

func (this *TuiCool) String() (rv string) {

	rv = fmt.Sprintf("title:%s\nlabel:%s\norigin:%s\ndate:%d\ncontent:%s\n", this.title, strings.Join(this.label, ","), this.origin, this.timestamp, this.content)
	return
}

func NewTuiCool() (rv *TuiCool) {
	rv = new(TuiCool)
	rv.label = make([]string, 0)
	return
}

func (this *TuiCool) ParseUrlToList(doc *goquery.Document) (list []string, err error) {
	article := regexp.MustCompile(`^/articles/\w+$`)
	doc.Find("a").Each(func(n int, g *goquery.Selection) {
		url, yes := g.Attr("href")
		if yes {
			if article.MatchString(url) {
				fmt.Println(url)
				list = append(list, url)
			}
			/*if strings.HasPrefix(url, "/articles/") {
				list = append(list, url)
			}*/
		}
	})

	return
}

func Save() {
}

func (this *TuiCool) GetContent(doc *goquery.Document) (err error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	article := doc.Find(".article_detail_bg")
	this.title = article.Find("h1").Text()
	article.Find(".new-label").Each(func(n int, g *goquery.Selection) {
		this.label = append(this.label, g.Text())
	})

	body, err := article.Find(".article_body").Html()
	this.content = strings.TrimSpace(body)

	date := strings.TrimSpace(strings.Split(article.Find(".timestamp").Text(), string(0xa0))[1])
	fmt.Println(date)
	p_t, err := time.Parse("2006-01-02 15:04:05", date)
	this.timestamp = p_t.Unix()
	this.origin = article.Find(".source").Find("a").Text()

	return
}
