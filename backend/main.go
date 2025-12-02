package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dralt03/AgentReport/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/scrape", api.ScrapeHandler).Methods("GET")
	r.HandleFunc("/items", api.ItemHandler).Methods("POST")

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	fmt.Println("Started Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", cors(r)))
}
