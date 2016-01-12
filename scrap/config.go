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

const KEYLOG_FORMAT = "========================== %s : %s ==========================\n"
const ITEMLOG_FORMAT = "%2d [%s<%s>]:%8s %s (%s)\n"
const SUNING = "苏宁易购"
const JD = "京东商城"

func ParseTitle(text string) string {
	title_rex, _ := regexp.Compile("\\s+|\\d+关注")
	return strings.TrimSpace(title_rex.ReplaceAllString(text, " "))
}

