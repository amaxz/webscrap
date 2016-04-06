package scrap

import (
	"log"
	"time"
	"net/http"
	httpclient "github.com/mreiferson/go-httpclient"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/url"
)

func Tmall(keyword string) ([]Item, string) {
	targeturl := "https://list.tmall.com/search_product.htm?q=" + url.QueryEscape(keyword) + "&cat=50024400&type=p&sort=d&spm=a220m.1000858.1000721.2.NpqWbx&from=.list.pc_1_searchbutton"
	transport := &httpclient.Transport{
		ConnectTimeout:        15 * time.Second,
		RequestTimeout:        20 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
	}
	defer transport.Close()
	client := &http.Client{Transport: transport}

	response, err := client.Do(newTmallRequest(targeturl))
	defer response.Body.Close()
	//utfBody, err := iconv.NewReader(response.Body, "gbk", "utf-8")
	utfBody := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Html())

	nodes := doc.Find(".product-iWrap")
	items := []Item{}
	nodes.Each(func(i int, s *goquery.Selection) {
		a := s.Find(".productTitle a")
		if a.Text() == "" { return }

		item := Item{}
		item.Title = ParseTitle(a.Text())
		item.Price = ParsePriceText(s.Find(".productPrice").Text())
		item.Vendor = ParseTitle(s.Find(".productShop").Text())

		if id, exists := s.Parent().Attr("data-id"); exists {
			item.Id = id
			item.Url = "http://detail.tmall.com/item.htm?id=" + id
		} else if href, exists := a.Attr("href"); exists {
			if strings.Index(href, "http") < 0 {href = "http:" + href}
			item.Url = href
		}
		items = append(items, item)
	})

	return items, targeturl
}

func newTmallRequest(targeturl string) *http.Request {
	request, _ := http.NewRequest("GET", targeturl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.7 (KHTML, like Gecko) Version/9.0.1	Safari/601.2.7")
	request.Header.Set("Cookie", "isg=2B2EC2C3E9E59A6222EBBB48A971E7A0; l=At/f7Abdbq/Ci4FoyQu9JKUJTz1pRDPm; cookie2=1cb86a2d9106fdd5f6c1b4613a671caf; t=6ff901036d21448d356bac63f39a2b75; _tb_token_=7NVVbAqTC9SA; pnm_cku822=213UW5TcyMNYQwiAiwTR3tCf0J%2FQnhEcUpkMmQ%3D%7CUm5OcktzTHhAekB5QHlGfyk%3D%7CU2xMHDJ%2BH2QJZwBxX39RaFZ4WHYxWDMdSx0%3D%7CVGhXd1llXGRbb1dtV25XblFoX2JAfUZ%2BQn9BekJ2THBLf0B0Wgw%3D%7CVWldfS0QMAo0Di4QMB41FSsFUwU%3D%7CVmhIGCQZJgY7GycYLRAwCzMMORklGi8SMgc6BycbJBEsDDYMMmQy%7CV25Tbk5zU2xMcEl1VWtTaUlwJg%3D%3D; res=scroll%3A1280*6528-client%3A1280*261-offset%3A1280*6528-screen%3A1280*800; cq=ccp%3D1")
	return request
}

func (v TmallFetcher) Load(task *Task) {
	items, url := Tmall(task.Keyword)
	task.Url = url
	task.Status = 200
	if length := len(items); length > 0 {
		utc := strconv.FormatInt(time.Now().Unix(), 10)
		for index := 0; index < length; index++ {
			items[index].Stamp = utc
			items[index].Keyword = task.Keyword
			items[index].Site = task.Src
			if items[index].Vendor == "" {
				items[index].Vendor = TMALL
			}
		}
		task.Items = items
	} else {
		task.Status = 400
	}
}

//"golang.org/x/text/transform"
//"golang.org/x/text/encoding/unicode"
//func ToUTF8(gbkstr string) string {
//	result, _, _ := transform.String(unicode.UTF8.NewEncoder(), gbkstr)
//	return result
//}
