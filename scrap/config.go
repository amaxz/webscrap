package scrap
import (
	"regexp"
	"strings"
)

type Item struct {
	id    string
	title string
	price string
	catentry string
	url string
	vendor string
}

const KEYLOG_FORMAT = "\n     ========================== %s : %s ==========================\n"
const ITEMLOG_FORMAT = "Item:%2d: %8s %s (%s)\n"
const SUNING = "苏宁易购"
const JD = "京东商城"

func ParseTitle(text string) string {
	title_rex, _ := regexp.Compile("\\s+")
	return strings.TrimSpace(title_rex.ReplaceAllString(text, " "))
}

func ParsePrice(text string) string {
	price_rex, _ := regexp.Compile(".*¥(\\d+\\.?\\d+).*")
	return strings.TrimSpace(price_rex.ReplaceAllString(text, "$1"))
}

