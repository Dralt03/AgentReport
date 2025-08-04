package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dralt03/AgentReport/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/scrape", api.ScrapeHandler).Methods("GET")

	fmt.Println("Started Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
