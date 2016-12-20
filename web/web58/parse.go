package web58

import (
	"fmt"
	"net/http"

	"os"

	"github.com/PuerkitoBio/goquery"
)

func GetHouseList(res *http.Response) (list []string, err error) {

	doc, err := goquery.NewDocumentFromResponse(res)
	con, err := doc.Html()
	fmt.Println(con)
	if err != nil {
		return
	}
	doc.Find(".qj-rentd").Each(func(n int, g *goquery.Selection) {
		url, yes := g.Find("a").Attr("href")
		fmt.Println(url)
		if yes {
			list = append(list, url)
		}
	})
	fmt.Println("list=%q", list)

	return
}

func GetBaseInfo() (err error) {
	var fd *os.File
	fd, err = os.Open("58house.html")
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(fd)
	if err != nil {
		return
	}
	//price
	price := doc.Find(".house-price").Text()
	fmt.Println(price)
	// update time
	updateTime := doc.Find(".title-right-info").Find("span").Eq(0).Text()
	fmt.Println(updateTime)
	// visit count
	visit := doc.Find(".title-right-info").Find("span").Eq(1).Text()
	fmt.Println(visit)
	// the image list
	var imgList []string
	doc.Find(".house-images-list").Find("img").Each(func(n int, g *goquery.Selection) {
		url, yes := g.Attr("lazy_src")
		if yes {
			imgList = append(imgList, url)
		}
	})
	fmt.Println(imgList)
	// house type
	houseType := doc.Find(".house-type").Text()
	fmt.Println(houseType)
	// court
	Country := doc.Find(".xiaoqu").Find("a").Eq(0).Text()
	Town := doc.Find(".xiaoqu").Find("a").Eq(1).Text()
	Street := doc.Find(".xiaoqu").Nodes[0].LastChild.Data

	fmt.Printf("%s-%s-%s\n", Country, Town, Street)
	// contact

	contact := doc.Find(".person-contact").Find("span").Eq(1).Text()
	fmt.Println(contact)
	// phone
	phone := doc.Find(".tel-num").Text()
	fmt.Println(phone)
	// description
	desc := doc.Find(".description-content").Text()
	fmt.Println(desc)
	// gps
	/*re.findall(`.*\"I\"\:6691\,\"V\":\"(.*?)\"\}.*`, a)
	re.findall(`.*\"I\"\:6692\,\"V\":\"(.*?)\"\}.*`, a)*/
	return
}
