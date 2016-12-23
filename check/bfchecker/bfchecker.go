package bfchecker

import (
	"errors"

	"github.com/konghui/bloomfilter"
	"github.com/konghui/gospider/check"
)

type BFChecker struct {
	bf *bloomfilter.BloomFilter
}

func New(args []interface{}) (checker check.Checker, err error) {
	size, yes := args[0].(uint64)

	if !yes {
		err = errors.New("invalid first size parameter, uint64 needed")
		return
	}

	var funList = make([]string, 0)
	for _, fun := range args[1:] {
		funName, yes := fun.(string)
		if !yes {
			err = errors.New("invalid hashfunction parameter, string needed")
			return
		}
		funList = append(funList, funName)
	}

	checker, err = NewBFChecker(size, funList)

	return
}

func NewBFChecker(size uint64, funList []string) (checker *BFChecker, err error) {
	checker = new(BFChecker)
	checker.bf, err = bloomfilter.NewBloomFilter(size, funList)
	return
}

func (this *BFChecker) HasVisited(url string) (visited bool, err error) {
	visited = this.bf.GetString(url)
	return
}

func (this *BFChecker) Visit(url string) (err error) {
	this.bf.SetString(url)
	return
}

func init() {
	check.RegisterPoolCheckMethod("bloomfileter", New)
}
