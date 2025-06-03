package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Wold")
}

func main(){
	http.HandleFunc("/", helloHandler)
	fmt.Println("Started Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}