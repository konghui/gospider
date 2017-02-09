package encoding

import "github.com/konghui/gospider/encoding/yaml"

// name: factory.go
// date: 2017-02-09 15:20:52
// author: konghui@live.cn

type Parser interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

func ParserFactory(name string) (parser Parser) {
	switch name {
	case "yaml":
		parser = yaml.NewYamlPaser()
	default:
		parser = nil
	}
	return
}
