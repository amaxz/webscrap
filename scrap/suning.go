package scrap

import (
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
	Rs     []Res  `json:"rs"`
}

type Res struct {
	CmmdtyCode string   `json:"cmmdtyCode"`
	CatentryId string   `json:"catentryId"`
	Price      string   `json:"price"`
	VendorName string   `json:"vendorName"`
	//SnPrice    string   `json:"snPrice"`
}

func Suning(keyword string) ([]Item, string) {

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
		return []Item{}, targeturl
	}

	nodes := doc.Find(".wrap")
	items := make([]Item, nodes.Length())
	nodes.Each(func(i int, s *goquery.Selection) {

		a := s.Find("a.sellPoint")
		vband := ParseTitle(a.Text())
		href, _ := a.Attr("href")
		item := Item{title: vband, price: "0", url: href}

		if classstr, ok := s.Find(".hidenInfo").Attr("datapro"); ok {
			regx, _ := regexp.Compile("\\d+")

			if properties := regx.FindAllString(classstr, -1); len(properties) >= 2 {
				commodityid, catentryid := properties[0], properties[1]
				if len(commodityid) < len(catentryid) {
					commodityid, catentryid = properties[1], properties[0]
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

	return QueryPrice(client, items), targeturl
}

func QueryPrice(client *http.Client, items []Item) []Item {
	var ids string
	for i, length := 0, len(items); i < length; i ++ {
		ids += items[i].id + "_" + items[i].catentry
		if (i + 1) >= 10 && (i + 1) % 10 == 0 {
			FetchPrice(client, ids, items)
			ids = ""
		} else if i < length {
			ids += ","
		}
	}
	if ids != "" {
		FetchPrice(client, ids, items)
	}
	return items
}

func FetchPrice(client *http.Client, ids string, items []Item) {
	priceurl := "http://ds.suning.cn/ds/general/" + ids + "-9063-1--1--getDataFromDsServer2.json"// "-9063-2-0000000000-1--.json"
	request := NewRequest(priceurl)
	//fmt.Println("Fetch", priceurl)
	response, err := client.Do(request)
	defer response.Body.Close()

	if (err == nil) {

		body, _ := ioutil.ReadAll(response.Body)

		//fmt.Println("Parse json", priceurl, string(body))
		dat := ParseJson(body)

		for j := 0; j < len(items); j++ {
			for i := 0; i < len(dat.Rs); i++ {
				if strings.EqualFold(items[j].id, dat.Rs[i].CmmdtyCode) {
					items[j].price = dat.Rs[i].Price
					items[j].vendor = dat.Rs[i].VendorName
				}
			}
		}
	}
}

func ParseJson(body []byte) *Price {
	var dat Price
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		log.Printf(err.Error())
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
	log.Printf(KEYLOG_FORMAT, SUNING, keyword)
	keyword = url.QueryEscape(keyword)
	items, url := Suning(keyword)
	if length := len(items); length > 0 {
		count := 0
		//var lowest Item
		for index := 0; index < length; index++ {
			item := items[index]
			if (item.price != "" && item.price != "0" && item.vendor == "") {
				count += 1
				log.Printf(ITEMLOG_FORMAT, count, item.price, item.title, item.url)
			}
			//if ((item.price))
		}
		if count == 0 {

		}
	} else {
		log.Println("No Item: ", url)
	}
}
