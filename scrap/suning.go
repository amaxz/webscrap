package scrap

import (
	"fmt"
	"log"
	"time"
	"net/http"
	httpclient "github.com/mreiferson/go-httpclient"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Price struct {
	Status int `json:"status"`
	Rs    []Res  `json:"rs"`
}

type Res struct {
	CmmdtyCode string   `json:"cmmdtyCode"`
	CatentryId string   `json:"catentryId"`
	Price      string   `json:"price"`
	SnPrice    string   `json:"snPrice"`
}

func Suning(keyword string) []Item {

	targeturl := "http://search.suning.com/" + keyword + "/&ci=20006&iy=-1"//"/&sc=0&ct=1&st=0"
	request := NewRequest(targeturl)
	transport := &httpclient.Transport{
		ConnectTimeout:        10 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
	}
	client := &http.Client{Transport: transport}

	response, err := client.Do(request)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Html())

	if response.StatusCode != 200 {
		log.Println(response.StatusCode, request.URL.String(), doc.Find("title").Text())
		return []Item{}
	}

	nodes := doc.Find(".wrap")
	items := make([]Item, nodes.Length())
	nodes.Each(func(i int, s *goquery.Selection) {

		vband := ParseTitle(s.Find("a.sellPoint").Text())
		item := Item{title: vband, price: "0"}

		if classstr, ok := s.Find(".hidenInfo").Attr("datapro"); ok {
			regx, _ := regexp.Compile("\\d+")
			properties := regx.FindAllString(classstr, -1)
			if (len(properties) >= 2) {
				commodityid := properties[0]
				catentryid := properties[1]
				if len(commodityid) < len(catentryid) {
					commodityid = properties[1]
					catentryid = properties[0]
				}
				zeros := 18 - len(commodityid)
				for times := 0; times < zeros; times ++ {
					commodityid = "0" + commodityid
				}

				item.id = commodityid
				item.catentry = catentryid
				//fmt.Println("Commodity ID: ", item.id, "Catentry ID: ", item.catentry)
			}
		}
		items[i] = item
	})

	return QueryPrice(client, items)
}

func QueryPrice(client *http.Client, items []Item) []Item {
	length := len(items)
	priceurl := ""
	for i := 0; i < length && i < 20; i ++ {
		priceurl += items[i].id + "_" + items[i].catentry
		if (i < length - 1) {
			priceurl += ","
		}
	}
	priceurl = "http://ds.suning.cn/ds/general/" + priceurl + "-9063-2-0000000000-1--.json"
	request := NewRequest(priceurl)

	response, err := client.Do(request)
	defer response.Body.Close()
	if (err == nil) {

		body, _ := ioutil.ReadAll(response.Body)

		fmt.Println("Parse json", priceurl, string(body))
		dat := ParseJson(body)

		for j := 0; j < len(items); j++ {
			for i := 0; i < len(dat.Rs); i++ {
				if strings.EqualFold(items[j].id, dat.Rs[i].CmmdtyCode) {
					items[j].price = dat.Rs[i].Price
					fmt.Printf(ITEMLOG_FORMAT, j, items[j].price, items[j].title)
				}
			}
		}
	}
	return items
}

func ParseJson(body []byte) *Price {
	var dat Price
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		fmt.Printf(err.Error())
	}
	return &dat
}

func NewRequest(targeturl string) *http.Request {
	request, _ := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	request.Header.Set("Cookie", "_snma=1%7C145216382589538017%7C1452163825895%7C1452191870882%7C1452191887535%7C12%7C2; _snsr=direct%7Cdirect%7C%7C%7C; _snmp=145219188671757070; __wmv=1452163826.2; SN_CITY=210_771_1000063_9063_01_10405_2_0; cityId=9063; districtId=10405; _customId=667644515806; _snmc=1; _snmb=145219154127538276%7C1452191887584%7C1452191887537%7C10; authId=si693E6DB63E48A9F24BA2F5D1D0F297F2; cart_abtest_num=39; cart_abtest=A; sesab=a; sesabv=35%23100%3A0; _snms=145219188758484108; _snck=14521925403882401")
	return request
}

func LoadSuning(keyword string) {
	fmt.Printf(KEYLOG_FORMAT, "Suning", keyword)
	keyword = url.QueryEscape(keyword)
	items := Suning(keyword)
	for index := 0; index < len(items); index++ {
		item := items[index]
		if (item.price != "" && item.price != "0") {
			log.Printf(ITEMLOG_FORMAT, index, item.price, item.title)
		}
	}
}
