package scraper

import (
	"encoding/json"
	"log"
	"strings"
	"unicode"

	"github.com/gocolly/colly/v2"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.Join(strings.Fields(text), " ")

	if len(text) > 0 {
		runes := []rune(text)
		runes[0] = unicode.ToUpper(runes[0])
		text = string(runes)
	}
	return text
}

func Scrape() string{
	c := colly.NewCollector(
		colly.AllowedDomains("cnn.com", "www.cnn.com"),
	)

	var articles []Article

	c.OnHTML("div.container__headline", func(e *colly.HTMLElement) {
		title := cleanText(e.Text)
		content := cleanText(e.DOM.Parent().Find(".container__description").Text())
		if title == "" {
			return
		}
		articles = append(articles, Article{
			Title:   title,
			Content: content,
		})
	})

	// Start scraping
	err := c.Visit("https://www.cnn.com")
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}