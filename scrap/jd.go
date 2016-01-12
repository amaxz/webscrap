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
)

func Jd(keyword string) ([]Item, string) {

	targeturl := "http://search.jd.com/Search?keyword=" + keyword + "&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&sttr=1&click=1&cid3=655&psort=2&stock=1&click=1&wtype=1"
	request, err := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	//request.Header.Set("Cookie", "__jdc=122270672; mx=0_X; xtest=3178.7099.b7a782741f667201b54880c925faec4b.b7a782741f667201b54880c925faec4b; ipLoc-djd=1-72-2819-0; __jda=122270672.1466835460.1452133618.1452144837.1452151685.3; __jdv=122270672|direct|-|none|-; __jdu=1466835460; ipLocation=%u5357%u5b81; __jdb=122270672.2.1466835460|3.1452151685")
	request.Header.Set("Cookie", "__jda=122270672.50445783.1452368009.1452368009.1452368009.1; __jdb=122270672.1.50445783|1.1452368009; __jdc=122270672; __jdv=122270672|direct|-|none|-; __jdu=50445783")

	transport := &httpclient.Transport{
		ConnectTimeout:        5 * time.Second,
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

	nodes := doc.Find(".gl-item")
	items := make([]Item, nodes.Length())
	nodes.Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		href, _ := a.Attr("href")
		vband := ParseTitle(a.Text())
		vprice := ParsePrice(s.Find(".p-price").Text())
		items[i] = Item{title: vband, price: vprice, url: href}
	})

	return items, targeturl
}

func ParsePrice(text string) string {
	price_rex, _ := regexp.Compile(".*¥(\\d+\\.?\\d+).*")
	return strings.TrimSpace(price_rex.ReplaceAllString(text, "$1"))
}

func LoadJD(keyword string) {
	log.Printf(KEYLOG_FORMAT, JD, keyword)
	key := url.QueryEscape(keyword)
	items, url := Jd(key)
	if length := len(items); length > 0 {
		for index := 0; index < length; index++ {
			item := items[index]
			log.Printf(ITEMLOG_FORMAT, index + 1, JD, keyword, item.price, item.title, "http:" + item.url)
		}
	} else {
		log.Println("No Item: ", url)
	}
}

//"golang.org/x/text/transform"
//"golang.org/x/text/encoding/unicode"
//utfBody, err := iconv.NewReader(response.Body, "GBK", "utf-8")
//func ToUTF8(gbkstr string) string {
//	result, _, _ := transform.String(unicode.UTF8.NewEncoder(), gbkstr)
//	return result
//}


