package task

import "github.com/konghui/gospider/urltool"

type ClientTask struct {
	Url *urltool.URL
}

func ValidUrl(task ClientTask) (rv bool) {
	return true
}
