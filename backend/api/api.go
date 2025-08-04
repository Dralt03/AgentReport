package api

import (
	"fmt"
	"net/http"

	"github.com/Dralt03/AgentReport/scraper"
)

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := scraper.Scrape()
	fmt.Fprintf(w, result)
}
