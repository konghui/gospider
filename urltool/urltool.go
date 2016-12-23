package urltool

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type URL struct {
	Domain string
	Port   int
	Uri    string
	Proto  string
}

var url = `(https?)://(.*?):?(\d+)?(/.*)`
var urlperttern = regexp.MustCompile(url)

func NewURL(url string) (u *URL, err error) {
	urls := urlperttern.FindStringSubmatch(url)

	u = &URL{Domain: urls[2]}

	u.Proto = urls[1]
	u.Uri = urls[4]

	if urls[3] == "" {
		if u.Proto == "http" {
			u.Port = 80
		} else if u.Proto == "https" {
			u.Port = 443
		} else {
			err = errors.New("unknown proto")
			return
		}
	} else {
		u.Port, err = strconv.Atoi(urls[3])
		if err != nil {
			return
		}
	}

	return
}

func (this *URL) String() (u string) {

	u = fmt.Sprintf("%s://%s:%d%s", this.Proto, this.Domain, this.Port, this.Uri)
	return
}

func (this *URL) StandUrl() (u string) {
	return this.String()
}
