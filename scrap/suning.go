package scrap

import (
	"fmt"
	"log"
	"time"
	"net/http"
	httpclient "github.com/mreiferson/go-httpclient"
	"github.com/PuerkitoBio/goquery"
//"golang.org/x/text/transform"
//"golang.org/x/text/encoding/unicode"
	"net/url"
)


func Suning(keyword string) []Item {

	targeturl := "http://search.suning.com/" + keyword + "/&sc=0&ct=1&st=0"
	request, err := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	request.Header.Set("Cookie", "_snma=1%7C145216382589538017%7C1452163825895%7C1452191870882%7C1452191887535%7C12%7C2; _snsr=direct%7Cdirect%7C%7C%7C; _snmp=145219188671757070; __wmv=1452163826.2; SN_CITY=210_771_1000063_9063_01_10405_2_0; cityId=9063; districtId=10405; _customId=667644515806; _snmc=1; _snmb=145219154127538276%7C1452191887584%7C1452191887537%7C10; authId=si693E6DB63E48A9F24BA2F5D1D0F297F2; cart_abtest_num=39; cart_abtest=A; sesab=a; sesabv=35%23100%3A0; _snms=145219188758484108; _snck=14521925403882401")

	transport := &httpclient.Transport{
		ConnectTimeout:        5 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
	}
	client := &http.Client{Transport: transport}

	response, err := client.Do(request)
	//	utfBody, err := iconv.NewReader(response.Body, "GBK", "utf-8")

	defer response.Body.Close()


	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Html())

	nodes := doc.Find(".wrap")
	items := make([]Item, nodes.Length())
	nodes.Each(func(i int, s *goquery.Selection) {
		vband := ParseTitle(s.Find("a").Text())
		vprice := ParsePrice(s.Find("strong").Text())
		fmt.Printf(ITEMLOG_FORMAT, i, vprice, vband)
		items[i] = Item{title: vband, price: vprice}
	})

	return items
}

func LoadSuning(keyword string) {
	fmt.Printf(KEYLOG_FORMAT, keyword)
	keyword = url.QueryEscape(keyword)
	items := Suning(keyword)
	for index := 0; index < len(items); index++ {
		item := items[index]
		log.Printf(ITEMLOG_FORMAT, index, item.price, item.title)
	}
}
//func ToUTF8(gbkstr string) string {
//	result, _, _ := transform.String(unicode.UTF8.NewEncoder(), gbkstr)
//	return result
//}


