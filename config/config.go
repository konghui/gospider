package config

// name: config.go
// date: 2017-02-09 13:48:00
// author: konghui@live.cn

import (
	"fmt"
	"io/ioutil"

	"github.com/konghui/gospider/encoding"
)

type Config struct {
	parser encoding.Parser
	args   interface{}
}

func NewConfig(parser encoding.Parser, in interface{}) (config *Config) {
	config = new(Config)
	config.parser = parser
	config.args = in
	return
}

func (this *Config) ReadFile(file string) (err error) {
	var buffer []byte

	buffer, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = this.parser.Unmarshal(buffer, this.args)
	return
}

func (this *Config) WriteFile(file string) (err error) {
	var buffer []byte

	buffer, err = this.parser.Marshal(this.args)
	err = ioutil.WriteFile(file, buffer, 0777)
	return
}

func (this *Config) show() {
	fmt.Println(this.args)
}
