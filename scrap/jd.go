package scrap

import (
	"log"
	"time"
	"net/http"
	httpclient "github.com/mreiferson/go-httpclient"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
	"regexp"
	"strconv"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"unicode/utf8"
)

type ShopInfo struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	VenderId string `json:"venderId,string"`
}

func Jd(keyword string) ([]Item, string) {
	targeturl := "http://search.jd.com/Search?keyword=" + url.QueryEscape(keyword) + "&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&sttr=1&click=1&cid3=655"//&psort=2&stock=1&click=1&wtype=1"
	transport := &httpclient.Transport{
		ConnectTimeout:        5 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
	}
	client := &http.Client{Transport: transport}

	response, err := client.Do(newJDRequest(targeturl))

	defer response.Body.Close()


	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	nodes := doc.Find(".gl-item")
	items := make([]Item, nodes.Length())
	var shopids = make(map [string] []string)
	nodes.Each(func(i int, s *goquery.Selection) {
		a := s.Find(".p-name a")
		items[i] = Item{}
		if id, exists := s.Attr("data-sku"); exists {
			items[i].Id = id
		}
		if href, exists := a.Attr("href"); exists {
			if strings.Index(href, "http") < 0 {href = "http:" + href}
			items[i].Url = href
		}
		if shopid, exists := s.Find(".p-shop").Attr("data-shopid"); exists {
			if shopid != "" {
				items[i].Catentry = shopid
				shopids[shopid] = append(shopids[shopid], items[i].Id)
			}
		}

		items[i].Title = ParseTitle(a.Text())
		items[i].Price = ParsePriceText(s.Find(".p-price").Text())
	})
	if len(shopids) > 0 {
		var ids []string
		for k, _ := range shopids {
			ids = append(ids, k)
		}
		fmt.Println(ids)
		items = FetchVender(client, strings.Join(ids, ","), items)
	}

	return items, targeturl
}

func FetchVender(client *http.Client, ids string, items []Item) []Item {
	vendorUrl := "http://search.jd.com/ShopName.php?ids=" + ids
	request := newJDRequest(vendorUrl)
	fmt.Println("Fetch", vendorUrl)
	response, err := client.Do(request)
	defer response.Body.Close()

	if (err == nil) {

		body, _ := ioutil.ReadAll(response.Body)

		fmt.Println("Parse json", vendorUrl, string(body))
		if _, start := utf8.DecodeRune(body); start > 0 {
			body = body[start:]
		}
		shops := make([]ShopInfo, 0)
		if err := json.Unmarshal(body, &shops); err != nil {
			fmt.Println(body)
			for i, l := 0, len(items); i < l; i++  {
				for _, shop := range shops  {
					if shop.Id == items[i].Catentry {
						items[i].Vendor = shop.Title
						/*title := ""
						for _, r := range []rune(shop.Title) {
							rint := int(r)
							if rint >= 128 {
								title += strconv.QuoteRune(r)
							} else {
								title += string(r)
							}
						}
						value.Vendor = strings.Replace(title, "'", "", -1)*/
					}
				}
			}
		}
	}
	return items
}

func newJDRequest(targeturl string) *http.Request {
	request, _ := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	//request.Header.Set("Cookie", "__jdc=122270672; mx=0_X; xtest=3178.7099.b7a782741f667201b54880c925faec4b.b7a782741f667201b54880c925faec4b; ipLoc-djd=1-72-2819-0; __jda=122270672.1466835460.1452133618.1452144837.1452151685.3; __jdv=122270672|direct|-|none|-; __jdu=1466835460; ipLocation=%u5357%u5b81; __jdb=122270672.2.1466835460|3.1452151685")
	request.Header.Set("Cookie", "__jda=122270672.50445783.1452368009.1452368009.1452368009.1; __jdb=122270672.1.50445783|1.1452368009; __jdc=122270672; __jdv=122270672|direct|-|none|-; __jdu=50445783")

	return request
}

func ParsePriceText(text string) string {
	price_rex, _ := regexp.Compile(".*¥(\\d+\\.?\\d+).*")
	return strings.TrimSpace(price_rex.ReplaceAllString(text, "$1"))
}

func (v JdFetcher) Load(task *Task) {
	items, url := Jd(task.Keyword)
	task.Items = items
	task.Url = url
	task.Status = 200
	if length := len(items); length > 0 {
		utc := strconv.FormatInt(time.Now().Unix(), 10)
		for index := 0; index < length; index++ {
			item := items[index]
			item.Stamp = utc
			item.Vendor = JD
		}
	} else {
		task.Status = 400
	}
}

//"golang.org/x/text/transform"
//"golang.org/x/text/encoding/unicode"
//utfBody, err := iconv.NewReader(response.Body, "GBK", "utf-8")
//func ToUTF8(gbkstr string) string {
//	result, _, _ := transform.String(unicode.UTF8.NewEncoder(), gbkstr)
//	return result
//}


