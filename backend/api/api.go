package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dralt03/AgentReport/scraper"
	"github.com/joho/godotenv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := scraper.Scrape()
	fmt.Fprintf(w, result)
}


func ItemHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var db *sql.DB
	if err := godotenv.Load(); err != nil {
        log.Println(".env not found")
    }

	connString := os.Getenv("DATABASE_URL")
	if(connString == ""){
		log.Print("DATABASE URL NOT FOUND")
	}

	db, err := sql.Open("postgres", connString)
    if err != nil {
        log.Fatal(err)
    }

	defer db.Close()

	rows, err := db.Query("SELECT id, title, content FROM Articles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Value); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, i)
	}

	json.NewEncoder(w).Encode(items)
}