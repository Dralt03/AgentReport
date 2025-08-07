package scraper

import (
	"regexp"
	"strings"
	"unicode"
)

var htmlTagRegex = regexp.MustCompile("<.*?>")

func stripHTML(s string) string {
    return htmlTagRegex.ReplaceAllString(s, "")
}

func cleanText(text string) string {

	text = stripHTML(text)
    
    text = strings.TrimSpace(text)
    text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	if len(text) > 0 {
		runes := []rune(text)
		runes[0] = unicode.ToUpper(runes[0])
		text = string(runes)
	}
	return text
}

