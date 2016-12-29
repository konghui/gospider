package web

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/konghui/gospider/check"
	"github.com/konghui/gospider/queue"
	"github.com/konghui/gospider/task"
)

type SpiderArgs struct {
	Doc      *goquery.Document
	Response *http.Response
	VisitMap *check.VisitPool
	Queue    *queue.Queue
	Task     *task.ClientTask
}
