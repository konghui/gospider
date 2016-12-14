package queue

import (
	"sync"

	log "github.com/Sirupsen/logrus"
)

type Queue struct {
	queue  []string
	mutex  sync.Mutex
	length uint32
}

var MyQueue = new(Queue)

func (this *Queue) Append(task string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.queue = append(this.queue, task)
	this.length = this.length + 1
	log.Debugf("add new task:%s to the queue", task)
}

func (this *Queue) Pop() (data string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.length == 0 {
		return ""
	}
	data = this.queue[0]
	this.queue = this.queue[1:this.length]
	this.length = this.length - 1
	log.Debugf("pop the task:%s from queue", data)
	return
}

func (this *Queue) Length() (len uint32) {
	return this.length
}
