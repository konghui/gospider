package config

// name: config_test.go
// date: 2017-02-09 15:44:20
// author: konghui@live.cn

import (
	"fmt"
	"testing"

	"github.com/konghui/gospider/encoding"
)

type demo struct {
	Name  string
	Debug string
}

func Test_ReadFile(t *testing.T) {
	var mydemo demo
	code := encoding.ParserFactory("yaml")
	if code == nil {
		t.Error("unknown encoding yaml\n")
	}
	config := NewConfig(code, &mydemo)
	err := config.ReadFile("config.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Print(mydemo)
}
