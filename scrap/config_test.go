package scrap

import (
	"testing"
	"strings"
	"fmt"
	"os"
	"bufio"
)

func TestParsePrice(t *testing.T) {
	text1 := ParsePrice("\rReview 0:   - 	¥1499.00\r			货到付款 ")
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

func TestFormatKey(t *testing.T)  {
	text1 := "三星 （SAMSUNG） G9200 FDD-LTE/TD-LTE版 32G, 金色"
	fmt.Println("Format before", text1)
	text1 = FormatKey(text1)
	fmt.Println("Format after", text1)
	if strings.Index(text1, "FDD") > 0 {
		t.Error("Formate error")
	}

	if file, err := os.OpenFile("/Volumes/Authoring/gopath/src/webscrap/input.txt", os.O_RDWR, os.ModePerm); err == nil {

		defer file.Close()

		re := bufio.NewReader(file)
		for i := 0; ; i ++ {
			line, err := re.ReadString('\n')
			if err != nil {
				break
			}
			if string(line) != "" {
				fmt.Println(line)
				fmt.Println("Text", FormatKey(line))
			}
		}
	} else {
		panic(err)
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