package web58

import "testing"

func Test_GetUrlList(t *testing.T) {
	var url = "http://bj.58.com/tiantongyuan/zufang/0/j2/"

	list, err := GetUrlList(url)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("get the house list:\n")
	for i, v := range list {
		t.Logf("%d:%s\n", i, v)
	}
	//	fmt.Println(list)
	//	t.Log(list)
}

func Test_GetBaseInfo(t *testing.T) {
	err := GetBaseInfo()
	if err != nil {
		t.Fatal(err.Error())
	}
}
