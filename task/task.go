package task

import "github.com/konghui/gospider/urltool"

type ClientTask struct {
	url *urltool.URL
}

func NewTask(url string) (task *ClientTask, err error) {
	task = new(ClientTask)
	task.url, err = urltool.NewURL(url)
	return
}

func (this *ClientTask) GetUrl() (url *urltool.URL) {
	url = this.url
	return
}

func ValidUrl(task ClientTask) (rv bool) {
	return true
}
