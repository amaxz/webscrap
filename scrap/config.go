package scrap
import (
	"regexp"
	"strings"
)

type Item struct {
	title string
	price string
}

const KEYLOG_FORMAT = "     ========================== %s ==========================\n"
const ITEMLOG_FORMAT = "Item:%2d: %8s %s\n"


func ParseTitle(text string) string {
	title_rex, _ := regexp.Compile("\\s+")
	return strings.TrimSpace(title_rex.ReplaceAllString(text, " "))
}

func ParsePrice(text string) string {
	price_rex, _ := regexp.Compile(".*¥(\\d+\\.?\\d+).*")
	return strings.TrimSpace(price_rex.ReplaceAllString(text, "$1"))
}

