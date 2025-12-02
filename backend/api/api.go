package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"io"
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

	body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }
	
	var payload map[string]interface{}
    if err := json.Unmarshal(body, &payload); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    toolCallId := ""
    if v, ok := payload["toolCallId"].(string); ok {
        toolCallId = v
    }

    if toolCallId == "" {
        log.Println("Missing toolCallId in incoming payload:", string(body))
        http.Error(w, "Missing toolCallId", http.StatusBadRequest)
        return
    }
	
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

	var sb strings.Builder
    for _, item := range items {
        sb.WriteString(fmt.Sprintf("• %s — %s\n", item.Name, item.Value))
    }

    response := map[string]interface{}{
        "results": []map[string]string{
            {
                "toolCallId": toolCallId,
                "result":     sb.String(),
            },
        },
    }

    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}