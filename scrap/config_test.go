package scrap

import (
	"testing"
	"strings"
	"fmt"
)

func TestParsePrice(t *testing.T) {
	text1 := ParsePrice("\rReview 0:   - 	¥1499.00\r			货到付款 ")
	if (strings.Index(text1, "1499.00") != 0) {
		t.Error("Price must equals 149900", text1)
	}
	fmt.Printf("a:b %8s ok\n", text1)

	text2 := ParseTitle("小米 4c 高配版 全网通 白色 移动联通电信4G手机 双卡双待 22695关注 - ")
	fmt.Printf("Item:%s %s OK\r", text2, text1)
	if (strings.Index(text2, " ") == 0) {
		t.Error("Title start with space char")
	}
}