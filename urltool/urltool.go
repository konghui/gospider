package urltool

import (
	"fmt"
	"regexp"
	"strconv"
)

type URL struct {
	Domain string
	Port   int
	Uri    string
	Ssl    bool
}

var url = `(https?)://(.*?):?(\d+)?(/.*)`
var urlperttern = regexp.MustCompile(url)

func NewURL(url string) (u *URL, err error) {
	urls := urlperttern.FindStringSubmatch(url)

	u = &URL{Domain: urls[2]}
	if urls[3] == "" {
		u.Port = 80
	} else {
		u.Port, err = strconv.Atoi(urls[3])
		if err != nil {
			return
		}
	}

	u.Ssl = urls[1] == "https"
	u.Uri = urls[4]

	return
}

func (this *URL) String() (u string) {

	var proto string

	if this.Ssl {
		proto = "https"
	} else {
		proto = "http"
	}

	u = fmt.Sprintf("%s://%s:%d%s", proto, this.Domain, this.Port, this.Uri)
	return
}

func (this *URL) StandUrl() (u string) {
	return this.String()
}
