package scrap

import (
	"fmt"
	"log"
	"time"
	"net/http"
	httpclient "github.com/mreiferson/go-httpclient"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

func JD(keyword string) []Item {

	targeturl := "http://search.jd.com/Search?keyword=" + keyword + "&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&sttr=1&click=1&cid3=655&psort=2&stock=1&click=1&wtype=1"
	request, err := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	request.Header.Set("Cookie", "__jdc=122270672; mx=0_X; xtest=3178.7099.b7a782741f667201b54880c925faec4b.b7a782741f667201b54880c925faec4b; ipLoc-djd=1-72-2819-0; __jda=122270672.1466835460.1452133618.1452144837.1452151685.3; __jdv=122270672|direct|-|none|-; __jdu=1466835460; ipLocation=%u5317%u4EAC; __jdb=122270672.2.1466835460|3.1452151685")

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
		vband := ParseTitle(s.Find("a").Text())
		vprice := ParsePrice(s.Find(".p-price").Text())
		fmt.Printf(ITEMLOG_FORMAT, i + 1, vprice, vband)
		items[i] = Item{title: vband, price: vprice}
	})

	return items
}


func LoadJD(keyword string) {
	fmt.Printf(KEYLOG_FORMAT, "京东商城", keyword)
	keyword = url.QueryEscape(keyword)
	items := JD(keyword)
	for index := 0; index < len(items); index++ {
		item := items[index]
		log.Printf(ITEMLOG_FORMAT, index + 1, item.price, item.title)
	}
}

//"golang.org/x/text/transform"
//"golang.org/x/text/encoding/unicode"
//utfBody, err := iconv.NewReader(response.Body, "GBK", "utf-8")
//func ToUTF8(gbkstr string) string {
//	result, _, _ := transform.String(unicode.UTF8.NewEncoder(), gbkstr)
//	return result
//}


