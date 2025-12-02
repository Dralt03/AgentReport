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
		log.Printf("Failed to read body: %v", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Debug: Log the entire payload to see what fields are present
	log.Printf("Received payload: %+v", payload)
	
	// Extract tool_call_id (Vapi uses snake_case, not camelCase)
	toolCallId := ""
	if v, ok := payload["tool_call_id"].(string); ok {
		toolCallId = v
	} else if v, ok := payload["toolCallId"].(string); ok {
		toolCallId = v
	}
	
	log.Printf("Extracted toolCallId: '%s'", toolCallId)
	
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Print("DATABASE URL NOT FOUND")
		response := map[string]interface{}{
			"results": []map[string]interface{}{
				{
					"toolCallId": toolCallId,
					"error":      "Database configuration error",
				},
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("Database connection error: %v", err)
		response := map[string]interface{}{
			"results": []map[string]interface{}{
				{
					"toolCallId": toolCallId,
					"error":      "Failed to connect to database",
				},
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content FROM Articles")
	if err != nil {
		log.Printf("Database query error: %v", err)
		response := map[string]interface{}{
			"results": []map[string]interface{}{
				{
					"toolCallId": toolCallId,
					"error":      fmt.Sprintf("Database query failed: %v", err),
				},
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Value); err != nil {
			log.Printf("Row scan error: %v", err)
			continue 
		}
		items = append(items, i)
	}

	// Build result string
	var resultText string
	if len(items) == 0 {
		resultText = "No news articles found in the database. The scraper may not have run yet."
	} else {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Found %d news articles:\n\n", len(items)))
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("• %s\n  %s\n\n", item.Name, item.Value))
		}
		resultText = sb.String()
	}

	response := map[string]interface{}{
		"results": []map[string]interface{}{
			{
				"toolCallId": toolCallId,
				"result":     resultText,
			},
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}