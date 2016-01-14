package scrap
import (
	"regexp"
	"strings"
	"encoding/json"
)

type Item struct {
	Id       string    `json:"-"`
	Stamp    string     `json:"stamp"`
	Vendor   string   `json:"vendor,omitempty"`
	Keyword  string  `json:"q"`
	Url      string      `json:"url"`
	Price    string    `json:"price"`
	Catentry string `json:"-"`
	Title    string    `json:"title"`
}

type Fetch struct {
	Keyword string
	Url     string
	Items   []Item
	Status  int
}

const KEYLOG_FORMAT = "========================== %s: %s ==========================\n"
const ITEMLOG_FORMAT = "%2d %s %8s <%s> (%s)\n"
const SUNING = "苏宁易购"
const JD = "京东商城"

func FormatKey(key string) string {

	regx3, _ := regexp.Compile("\\s\\d\\.\\d+\\s|FDD-LTE/TD-LTE版|FDD-LTE版|TD-LTE版|TD-SCDMA版|WCDMA版|GSM版")
	key = regx3.ReplaceAllString(key, "")

	regx1, _ := regexp.Compile("（.*）|/|,|\\+|\\.")
	key = regx1.ReplaceAllString(key, " ")

	regx2, _ := regexp.Compile("\\s\\s+")
	key = regx2.ReplaceAllString(key, " ")

	return key
}

func ParseTitle(text string) string {
	title_rex, _ := regexp.Compile("\\s+|\\d+关注")
	return strings.TrimSpace(title_rex.ReplaceAllString(text, " "))
}

func JsonString(item interface{}) string {
	b, _ := json.Marshal(item)
	s := string(b)
	return s
}

