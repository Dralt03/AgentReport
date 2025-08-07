package scraper

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

//For news channels because Colly only scrapes raw HTML
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

func Scrape() string{
	
	var articles []Article
	scrapeSite(colly.NewCollector(colly.AllowedDomains("cnn.com", "edition.cnn.com", "www.cnn.com")), "cnn", &articles)
	scrapeSite(nil, "bbc", &articles)
	scrapeSite(nil, "cnbc", &articles)
	scrapeSite(nil, "guardian", &articles)

	
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}

var malformedFormat = regexp.MustCompile(`%[^\w]`)

func fetchRSS(url string, src string, articles *[]Article) {
    resp, err := http.Get(url)
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
		cleanTitle := cleanText(item.Title)
        cleanDesc := cleanText(item.Description)

		content := cleanText(strings.ReplaceAll(item.Description, "\"", "'"))

		if malformedFormat.MatchString(cleanTitle) || malformedFormat.MatchString(content) {
			continue
		}

		if !json.Valid([]byte(`"` + cleanTitle + `"`)) || !json.Valid([]byte(`"` + content + `"`)) {
			continue
		}

		cleanDesc = strings.ReplaceAll(cleanDesc, "\"", "'")
		
        *articles = append(*articles, Article{
            Title:    cleanTitle,
            Content:  cleanDesc,
            Src:      src,
        })
    }
}

func scrapeSite(c *colly.Collector, site string, articles *[]Article) {
	
	switch site {
	case "cnn":
		if c == nil {
			log.Fatal("Collector required for CNN")
		}
		c.OnHTML("div.container__headline", func(e *colly.HTMLElement) {
			title := cleanText(e.Text)
			content := cleanText(e.DOM.Parent().Find(".container__description").Text())
			
			if title != ""  && content != ""{
				*articles = append(*articles, Article{Title: title, Content: content, Src:"cnn"})
			}
		})
		err := c.Visit("https://www.cnn.com")
		if err != nil {
			log.Fatal(err)
		}
		

	case "bbc":
        fetchRSS("https://feeds.bbci.co.uk/news/rss.xml", "bbc", articles)

    case "cnbc":
        fetchRSS("https://www.cnbc.com/id/100003114/device/rss/rss.html", "cnbc", articles)

	case "guardian":
    	fetchRSS("https://www.theguardian.com/world/rss", "guardian", articles)
	default:
		log.Fatalf("Unsupported site: %s", site)
	}
}