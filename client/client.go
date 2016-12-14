package client

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/konghui/gospider/task"
	"github.com/konghui/gospider/urltool"
	"github.com/konghui/gospider/web"
)

type Client struct {
	httpClient *http.Client
	Site       map[string]*web.WebSite
}

func NewClient() (client *Client) {
	client = &Client{
		httpClient: &http.Client{},
		Site:       make(map[string]*web.WebSite),
	}
	return
}

func (this *Client) ConsumerLoop() {
	log.Info("Enter the consumer loop")
	for {
		var task task.ClientTask
		var err error
		var out string

		//url := GetHttpUrlList()
		// for test only modify it later
		var url string
		if url == "" {
			log.Info("url set is empty!")
			time.Sleep(1 * time.Second)
			continue
		}
		task.Url, err = urltool.NewURL(url)
		if err != nil {
			log.Warning(err.Error())
		}
		if web, yes := this.Site[task.Url.Domain]; yes == true {
			content, err := web.SendHttpRequest(&task)
			if err != nil {
				log.Warning(err.Error())
				continue
			}

			out, err = web.GetResponseContent(content)
			if err != nil {
				log.Warning(err.Error())
			}
			for _, v := range web.GetFilterRule() {

				groups := v.Parttern.FindStringSubmatch(out)
				fmt.Printf("%q", groups)
				if len(groups) < 2 {
					log.Warn("No match found for the rule:%s, page:%s", v.Rule, task.Url.StandUrl())
					continue
				}

				v.Data = groups[1]
				fmt.Println(v.Data)
			}
		}

	}
}
