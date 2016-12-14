package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type House struct {
	Price          int
	Class          string
	Desc           string
	Floor          string
	Court          string
	Infrastructure string
	Bref           string
	Address        *Addr
	Contact        *Contact
	Picture        []string
}

type Addr struct {
	City    string
	Country string
	Town    string
	Street  string
}

type Contact struct {
	Name  string
	Phone uint64
	Type  byte
}

func (this *Contact) String() (s string) {
	s = fmt.Sprintf("Name: %s\nType: %d\nPhone: %d", this.Name, this.Type, this.Phone)
	return
}

func (this *Addr) String() (s string) {
	s = fmt.Sprintf("City: %s\nCountry: %s\nTown: %s\nStreet: %s\n", this.City, this.Country, this.Town, this.Street)
	return
}
func (this *House) String() (s string) {
	s = fmt.Sprintf("Price: %d\nClass: %s\nBref: %s\nFloor: %s\nCourt: %s\nInfrastructure: %s\nDescription: %s\nPicture: %s\n%s\n%s\n", this.Price, this.Class, this.Bref, this.Floor, this.Court, this.Infrastructure, this.Desc, this.Picture, this.Address.String(), this.Contact.String())
	return
}

func main() {
	var myhouse House
	/*rent_math := `http://bj.58.com/zufang/.*\.shtml`
	price_match := `<b class="basic-info-price fl">\d*</b>`
	class := `"\s(.*)\s-\s(.*)\s-\s(.*)\s㎡`
	desc := `"\s(.*)\s-\s(.*)\s-\s(.*)\s"`
	floor := `"\s(.*)\s(\(.*\()\s"`
	zone := `<`
	out, _ := ioutil.ReadFile("test.html")
	parttern := regexp.MustCompile(rent_math)
	fmt.Printf("%q", parttern.FindAllStringSubmatch(string(out), 100))*/
	fd, err := os.Open("ganji.htm")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer fd.Close()
	doc, err := goquery.NewDocumentFromReader(fd)
	if err != nil {
		fmt.Println(err.Error())
	}

	myhouse.GetHouseInfo(doc)
	fmt.Println(myhouse.String())
}

func (this *House) GetHouseInfo(doc *goquery.Document) (err error) {
	info := doc.Find(".basic-info-ul").Find("li")
	this.Price, err = strconv.Atoi(info.Eq(0).Find(".basic-info-price").Text())
	if err != nil {
		return
	}
	this.Class = strings.TrimSpace(info.Eq(1).Nodes[0].LastChild.Data)
	this.Desc = strings.TrimSpace(info.Eq(2).Nodes[0].LastChild.Data)
	this.Floor = strings.TrimSpace(info.Eq(3).Nodes[0].LastChild.Data)
	//info.Find("li").Each(func(n int, t *goquery.Selection) { fmt.Printf("%d--%s\n", n, t.Text()) })
	this.Court = strings.TrimSpace(info.Eq(4).Find("a").Eq(0).Text())

	var addr Addr
	addrInfo := info.Eq(5).Find("a")

	addr.City = strings.TrimSpace(addrInfo.Eq(0).Text())
	addr.Country = strings.TrimSpace(addrInfo.Eq(1).Text())
	addr.Town = strings.TrimSpace(addrInfo.Eq(2).Text())
	addr.Street = strings.TrimSpace(info.Eq(6).Find(".addr-area").Text())

	this.Address = &addr
	this.Infrastructure = strings.TrimSpace(info.Eq(7).Find("p").Text())

	var contact Contact

	contactInfo := doc.Find(".basic-info-contact")
	nameAndType := contactInfo.Find("i")
	contact.Name = nameAndType.Eq(0).Text()
	Type := nameAndType.Eq(1).Text()
	fmt.Println(Type)
	if strings.Contains(nameAndType.Eq(1).Text(), "经纪人") {
		contact.Type = 1
	} else {
		contact.Type = 0
	}
	// remove the space between the phon number
	phone := strings.Replace(contactInfo.Find(".contact-mobile").Text(), " ", "", 2)

	contact.Phone, err = strconv.ParseUint(phone, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.Contact = &contact

	this.Desc = doc.Find(".summary-cont").Text()

	doc.Find(".pics").Find("img").Each(
		func(n int, g *goquery.Selection) {
			v, yes := g.Attr("src")
			if yes {
				this.Picture = append(this.Picture, v)
			}
		})

	data, yes := doc.Find("#map_load").Attr("data-ref")
	if yes {
		fmt.Println(data)
	}

	return
}
