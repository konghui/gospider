package main

import (
	"fmt"

	"github.com/konghui/gospider/encoding"
)

// name: setup.go
// date: 2017-02-09 16:39:52
// author: konghui@live.cn

type Args struct {
	Log_level string
}

func (this *Args) String() (s string) {
	s = fmt.Sprintf("log_level => %s\n", this.Log_level)
	return
}

var logType = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
}

func setupLog(level string) {
	v, ok = logType[level]
	if !ok {
		v = logType["debug"]
	}
	log.SetLevel(v)
}

func parseConfig(args *Args) (err error) {

	coder := encoding.ParserFactory("yaml")
	if code == nil {
		t.Error("unknown encoding yaml\n")
	}
	config := NewConfig(coder, args)
	err = config.ReadFile("config.yaml")

	return
}

func setup() (err error) {
	var args Args
	err = parseConfig(&args)
	if err != nil {
		return
	}
	setupLog(args.Log_level)

	return
}
