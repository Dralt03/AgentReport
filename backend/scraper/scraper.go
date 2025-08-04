package scraper

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/gocolly/colly/v2"
)

//For BBC because Colly only scrapes raw HTML
type RSS struct {
	Channel struct {
		Items []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Src string `json:"src"`
	TimeSpan string `json:"timeSpan"`
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
	// scrapeSite(c, "cnn", &articles)
	scrapeSite(c, "bbc", &articles)

	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}

func scrapeSite(c *colly.Collector, site string, articles *[]Article) {
	
	switch site {
	case "cnn":
		c.AllowedDomains = []string{"cnn.com", "www.cnn.com"}
		c.OnHTML("div.container__headline", func(e *colly.HTMLElement) {
			title := cleanText(e.Text)
			content := cleanText(e.DOM.Parent().Find(".container__description").Text())
			
			if title != "" {
				*articles = append(*articles, Article{Title: title, Content: content, Src:"cnn"})
			}
		})
		err := c.Visit("https://www.cnn.com")
		if err != nil {
			log.Fatal(err)
		}

	case "bbc":
		resp, err := http.Get("https://feeds.bbci.co.uk/news/rss.xml")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		var rss RSS
		if err := xml.Unmarshal(data, &rss); err != nil {
			log.Fatal(err)
		}

		
		for _, item := range rss.Channel.Items {
			if(item.Title == ""){
				continue
			}

			*articles = append(*articles, Article{
				Title:    cleanText(item.Title),
				Content:  cleanText(item.Description),
				Src:      "bbc",
				TimeSpan: strings.TrimSpace(item.PubDate),
			})
		}

	default:
		log.Fatalf("Unsupported site: %s", site)
	}
}