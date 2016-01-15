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
	"strconv"
)

type Price struct {
	Status int    `json:"status"`
	Rs     []Res  `json:"rs"`
}

type Res struct {
	CmmdtyCode string   `json:"cmmdtyCode"`
	CatentryId string   `json:"catentryId"`
	Price      string   `json:"price"`
	VendorName string   `json:"vendorName"`
}

func (v SuningFetcher) Suning(keyword string) ([]Item, string) {
	targeturl := "http://search.suning.com/" + url.QueryEscape(keyword) + "/&ci=20006&iy=-1"//"/&sc=0&ct=1&st=0"
	request := newRequest(targeturl)
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

	if items := make([]Item, nodes.Length()); len(items) > 0 {

		nodes.Each(func(i int, s *goquery.Selection) {

			a := s.Find("a.sellPoint")
			vband := ParseTitle(a.Text())
			href, _ := a.Attr("href")
			item := Item{Title: vband, Price: "0", Url: href, Keyword: keyword}

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

					item.Id = commodityid
					item.Catentry = catentryid
					//fmt.Println("Commodity ID: ", item.id, "Catentry ID: ", item.catentry)
				}
			}
			items[i] = item
		})
		return queryPrice(client, items), targeturl
	} else {
		log.Println("400", keyword, targeturl)
		return []Item{}, targeturl
	}
}

func queryPrice(client *http.Client, items []Item) []Item {
	var ids string
	for i, length := 0, len(items); i < length; i ++ {
		ids += items[i].Id + "_" + items[i].Catentry
		if (i + 1) >= 10 && (i + 1) % 10 == 0 {
			FetchPrice(client, ids, items)
			ids = ""
		} else if i < length - 1 {
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
	request := newRequest(priceurl)
	//fmt.Println("Fetch", priceurl)
	response, err := client.Do(request)
	defer response.Body.Close()

	if (err == nil) {

		body, _ := ioutil.ReadAll(response.Body)

		//fmt.Println("Parse json", priceurl, string(body))
		dat := parsePriceJson(body)

		for j := 0; j < len(items); j++ {
			for i := 0; i < len(dat.Rs); i++ {
				if strings.EqualFold(items[j].Id, dat.Rs[i].CmmdtyCode) {
					items[j].Price = dat.Rs[i].Price
					items[j].Vendor = dat.Rs[i].VendorName
				}
			}
		}
	}
}

func parsePriceJson(body []byte) *Price {
	var dat Price
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Printf(err.Error())
	}
	return &dat
}

func newRequest(targeturl string) *http.Request {
	request, _ := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	request.Header.Set("Cookie", "_snma=1%7C145216382589538017%7C1452163825895%7C1452191870882%7C1452191887535%7C12%7C2; _snsr=direct%7Cdirect%7C%7C%7C; _snmp=145219188671757070; __wmv=1452163826.2; SN_CITY=210_771_1000063_9063_01_10405_2_0; cityId=9063; districtId=10405; _customId=667644515806; _snmc=1; _snmb=145219154127538276%7C1452191887584%7C1452191887537%7C10; authId=si693E6DB63E48A9F24BA2F5D1D0F297F2; cart_abtest_num=39; cart_abtest=A; sesab=a; sesabv=35%23100%3A0; _snms=145219188758484108; _snck=14521925403882401")
	return request
}

func (v SuningFetcher) Load(task *Task) {
	items, url := v.Suning(task.Keyword)
	task.Items = items
	task.Url = url
	task.Status = 200
	if length := len(items); length > 0 {
		//count := 0
		//lowest_pos := -1
		//lowest_price := math.MaxFloat32
		utc := strconv.FormatInt(time.Now().Unix(), 10)
		for index := 0; index < length; index++ {
			item := items[index]
			item.Stamp = utc
			item.Vendor = SUNING
			//if (item.Price != "" && item.Price != "0" && item.Vendor == "") {
			//   count += 1
			//}
			//if p, err := strconv.ParseFloat(item.Price, 32); err == nil && p < lowest_price {
			//	lowest_pos = index
			//	lowest_price = p
			//}
		}
		/*if count == 0 && lowest_pos != -1 {
			items[lowest_pos].Stamp = utc
			fmt.Printf(ITEMLOG_FORMAT, lowest_pos + 1, SUNING, items[lowest_pos].Price, items[lowest_pos].Title, items[lowest_pos].Url)
			log.Println(JsonString(items[lowest_pos]))
		}*/
	} else {
		task.Status = 404
	}
}
