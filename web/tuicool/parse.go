package tuicool

import (
	"fmt"
	"strings"
	"time"

	"regexp"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/web"
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
	// callbak handler
	handler []func(*web.SpiderArgs) error
}

var list = make([]string, 0)

func (this *TuiCool) String() (rv string) {

	rv = fmt.Sprintf("title:%s\nlabel:%s\norigin:%s\ndate:%d\ncontent:%s\n", this.title, strings.Join(this.label, ","), this.origin, this.timestamp, this.content)
	return
}

func NewTuiCool() (rv *TuiCool) {
	rv = new(TuiCool)
	rv.label = make([]string, 0)
	rv.handler = make([]func(*web.SpiderArgs) error, 0)
	rv.handler = []func(*web.SpiderArgs) error{rv.ParseUrlToList, rv.GetContent}
	return
}

func (this *TuiCool) GetHandler() (fun []func(*web.SpiderArgs) error) {
	fun = this.handler
	return
}

func (this *TuiCool) ParseUrlToList(args *web.SpiderArgs) (err error) {
	article := regexp.MustCompile(`^/articles/\w+$`)
	args.Doc.Find("a").Each(func(n int, g *goquery.Selection) {
		url, yes := g.Attr("href")
		if yes {
			if article.MatchString(url) {
				list = append(list, url)
			}
		}
	})
	fmt.Println(list)
	for _, v := range list {
		newUrl := web.GetURLOfTheTaskPageLink(args.Task, v)
		if visited, err := args.VisitMap.HasVisited(newUrl); err == nil {
			if visited {
				log.Infof("url:%s already visit. skip it", newUrl)
				continue
			}
		} else {
			log.Warn(err.Error())
		}
		args.Queue.Append(newUrl)
		args.VisitMap.Visit(newUrl)
	}
	return
}

func Save() {
}

func (this *TuiCool) GetContent(args *web.SpiderArgs) (err error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	article := args.Doc.Find(".article_detail_bg")
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
	fmt.Println(this)
	return
}
