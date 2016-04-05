package scrap

import (
	"testing"
	"strings"
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	"strconv"
	"time"
)

func TestParsePrice(t *testing.T) {
	text1 := ParsePriceText("\rReview 0:   - 	¥1499.00\r			货到付款 ")
	if strings.Index(text1, "1499.00") != 0 {
		t.Error("Price must equals 149900", text1)
	}
	fmt.Printf("a:b %8s ok\n", text1)

	text2 := ParseTitle("小米 4c 高配版 全网通 白色 移动联通电信4G手机 双卡双待 22695关注 - ")
	fmt.Printf("Item:%s %s OK\r", text2, text1)
	if strings.Index(text2, " ") == 0 {
		t.Error("Title start with space char")
	}
}

func TestFormatKey(t *testing.T) {
	text1 := "三星 （SAMSUNG） G9200 FDD-LTE/TD-LTE版 32G, 金色"
	fmt.Println("Format before", text1)
	text1, _ = FormatKey(text1)
	fmt.Println("Format after", text1)
	if strings.Index(text1, "FDD") > 0 {
		t.Error("Formate error")
	}

	if file, err := os.OpenFile("/Volumes/Authoring/gopath/src/webscrap/input.txt", os.O_RDWR, os.ModePerm); err == nil {

		defer file.Close()

		re := bufio.NewReader(file)
		for i := 0;; i ++ {
			line, err := re.ReadString('\n')
			if err != nil {
				break
			}
			if string(line) != "" {
				fmt.Println(line)
				k, _ := FormatKey(line)
				fmt.Println("Text", k)
			}
		}
	} else {
		panic(err)
	}
}

func TestMap(t *testing.T) {
	shopids := make(map[string][]string)
	//s1 := make([]string, 5)
	//s2 := make([]string, 5)

	//shopids["aaa"] = &s1
	//shopids["aab"] = &s2

	shopids["aaa"] = append(shopids["aaa"], "ccc")
	shopids["aab"] = append(shopids["aab"], "ccc", "ddd")

	for k, v := range shopids {
		fmt.Println(k, v)
	}
}

func TestJson(t *testing.T) {
	item := Item{Id: "1", Price: "900.00", Vendor: "JD"}
	json := JsonString(item)
	if (strings.Index(json, "\"900.00\"") <= 0) {
		t.Error("Json Marshal Error")
	}
	if (strings.Index(json, "\"catentry\"") > 0) {
		t.Error("Json Marshal Error")
	}
	if (strings.Index(json, "\"id\"") > 0) {
		t.Error("Json Marshal Error")
	}
	fmt.Println(json)
}

func TestTimeFormat(t *testing.T) {
	fmt.Println(time.RFC850)
}

func TestFetchVender(t *testing.T) {

	//text := `[{"id":"160554","title":"\u7d22\u5c3c\u624b\u673a\u65d7\u8230\u5e97","url":"http:\/\/sonymobile.jd.com","venderId":165367}]`
	text := `[{"id":"1000004065","title":"OPPO\u624b\u673a\u5b98\u65b9\u65d7\u8230\u5e97","url":"http:\/\/oppo.jd.com","venderId":1000004065}]`

	dat := make([]ShopInfo, 0)
	if err := json.Unmarshal([]byte(text), &dat); err == nil {
		fmt.Println([]byte(text))
		for i, v := range dat {
			fmt.Println(i, v.Id)
			title := ""
			for _, r := range []rune(v.Title) {
				rint := int(r)
				if rint >= 128 {
					title += strconv.QuoteRune(r)
				} else {
					title += string(r)
				}
			}
			fmt.Println(strings.Replace(title, "'", "", -1))
		}
	} else {
		fmt.Println(err)
	}

	str := `OPPO\u624b\u673a\u5b98\u65b9\u65d7\u8230\u5e97`
	str, _ = strconv.Unquote("\"" + str + "\"")
	fmt.Println(str)


}