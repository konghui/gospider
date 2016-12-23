package check

import (
	"errors"

	"fmt"

	"github.com/konghui/gospider/urltool"
)

type Checker interface {
	//New(string, ...interface{})
	HasVisited(string) (bool, error)
	Visit(string) error
}

type VisitPool struct {
	checker Checker
}

func NewVisitPool(poolType string, args []interface{}) (pool *VisitPool, err error) {
	pool = new(VisitPool)
	initFunc, err := GetPoolCheckMethod(poolType)
	if err != nil {
		return
	}
	pool.checker, err = initFunc(args)
	if err != nil {
		return
	}
	return
}

var checkMethodDict = make(map[string]func([]interface{}) (Checker, error))

func RegisterPoolCheckMethod(name string, handler func([]interface{}) (Checker, error)) {
	checkMethodDict[name] = handler
}

func GetPoolCheckMethod(name string) (fun func([]interface{}) (Checker, error), err error) {
	var yes bool
	if fun, yes = checkMethodDict[name]; !yes {
		err = errors.New(fmt.Sprintf("Unknown pool type %s\n", name))
		return
	}
	return
}

func (this *VisitPool) HasVisited(url string) (visited bool, err error) {
	var standUrl *urltool.URL

	standUrl, err = urltool.NewURL(url)
	if err != nil {
		return
	}
	visited, err = this.checker.HasVisited(standUrl.String())
	return
}

func (this *VisitPool) Visit(url string) (err error) {
	var standUrl *urltool.URL

	standUrl, err = urltool.NewURL(url)
	if err != nil {
		return
	}
	err = this.checker.Visit(standUrl.String())
	return
}
