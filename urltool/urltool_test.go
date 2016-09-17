package urltool

import (
	"fmt"
	"testing"
)

func NewUrlTest(t *testing.T, url string) {
	u, err := NewURL(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log(u)
}

func Test_HttpUrlNoPort(t *testing.T) {
	NewUrlTest(t, "http://www.google.com/test1/2.html?login=aaa")
}

func Test_HttpsUrlNoPort(t *testing.T) {
	NewUrlTest(t, "https://www.google.com/test1/2.html?login=aaa")
}

func Test_HttpsUrlWithPort(t *testing.T) {
	NewUrlTest(t, "https://www.google.com:80/test1/2.html?login=aaa")
}

func Test_RootUrl(t *testing.T) {
	NewUrlTest(t, "https://www.google.com:80/")
}
