package yaml

// name: yaml.go
// date: 2017-02-09 15:02:32
// author: konghui@live.cn

import (
	"gopkg.in/yaml.v2"
)

type YamlParser struct {
}

func NewYamlPaser() (parser *YamlParser) {
	parser = new(YamlParser)
	return
}

func (this *YamlParser) Marshal(in interface{}) (buffer []byte, err error) {
	buffer, err = yaml.Marshal(in)
	return
}

func (this *YamlParser) Unmarshal(buffer []byte, out interface{}) (err error) {
	err = yaml.Unmarshal(buffer, out)
	return
}
