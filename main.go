package main

import (
	"fmt"
	"log"
	"net/http"
)

func proccessReceipt(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	points := calculatePoints(receipt)
	
	id := uuid.New().String()
	receipts[id] =  points

	json.NewEncoder(w).Encode(map[string]string{"id": id})
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

