package main

import (
	"fmt"
	"log"
	"net/http"
)

func proccessReceipt(w http.ResponseWriter, r *http.Request) {
	
}
func main() {

	http.HandleFunc("/receipts/proccess", proccessReceipt)
	http.Handle("/receipts", getPoints)

	fmt.Printf("Starting server at port 8080\n")

	err := http.ListenAndServe(":8080", nil); 
	if err != nil {
		log.Fatal(err)
	}
}

