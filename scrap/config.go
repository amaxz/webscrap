package scrap
import (
	"regexp"
	"strings"
	"encoding/json"
)

type Item struct {
	Id       string   `json:"-"`
	Site     string   `json:"src"`
	Stamp    string   `json:"stamp"`
	Keyword  string   `json:"q"`
	Url      string   `json:"url"`
	Vendor   string   `json:"vendor,omitempty"`
	Price    string   `json:"price"`
	Catentry string   `json:"-"`
	Title    string   `json:"title"`
}

type Fetcher interface {
	Load(task *Task)
}

type Task struct {
	Keyword string
	Url     string
	Items   []Item
	Status  int
	Src     string
	Fetcher Fetcher
	Brand   string
}

type TmallFetcher struct {
}

type JdFetcher struct {
}

type SuningFetcher struct {
}

const KEYLOG_FORMAT = "\n========================== %s: %s ==========================\n"
const ITEMLOG_FORMAT = "%2d %8s %s <%s> (%s)\n"
const SUNING = "苏宁易购"
const JD = "京东商城"
const TMALL = "天猫商城"

func FormatKey(key string) (string, string) {

	regx3, _ := regexp.Compile("\\s\\d\\.\\d+\\s|FDD-LTE/TD-LTE版|FDD-LTE版|TD-LTE版|TD-SCDMA版|WCDMA版|GSM版")
	key = regx3.ReplaceAllString(key, "")

	regx1, _ := regexp.Compile("（.*）|/|,|\\+|\\.")
	key = regx1.ReplaceAllString(key, " ")

	regx2, _ := regexp.Compile("\\s\\s+")
	key = regx2.ReplaceAllString(key, " ")

	return strings.TrimSpace(key), strings.Split(key, " ")[0]
}

func ParseTitle(text string) string {
	title_rex, _ := regexp.Compile("\\s+")
	return strings.TrimSpace(title_rex.ReplaceAllString(text, " "))
}

func JsonString(item interface{}) string {
	b, _ := json.Marshal(item)
	return string(b)
}

