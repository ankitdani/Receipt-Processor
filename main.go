package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "time"
    "github.com/google/uuid"
)

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}

type Receipt struct {
    Retailer     string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Items        []Item `json:"items"`
    Total        string `json:"total"`
}

type PointsResponse struct {
    Points int `json:"points"`
}

var receipts = make(map[string]int)

func calculatePoints(receipt Receipt) int {
    points := 0

    alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
    points += len(alphanumeric.FindAllString(receipt.Retailer, -1))

    total, err := strconv.ParseFloat(receipt.Total, 64)
    if err != nil {
        return 0, fmt.Errorf("Invalid total amount: %s", err)
    }

    if total == math.Floor(total) {
        points += 50
    }

    if int(total*100)%25 == 0 {
        points += 25
    }

    points += (len(receipt.Items) / 2) * 5

    for _, item := range receipt.Items {
        description := strings.TrimSpace(item.ShortDescription)
        if len(description)%3 == 0 {
            price, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(price * 0.2))
        }
    }

    day, err := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2])
    if err != nil {
        return 0, fmt.Errorf("Invalid purchase date: %s", err)
    }
    if day%2 != 0 {
        points += 6
    }

    purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
    if err != nil {
        return 0, fmt.Errorf("Invalid purchase time: %s", err)
    }
    if purchaseTime.After(time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC)) &&
        purchaseTime.Before(time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)) {
        points += 10
    }

    return points
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt Receipt

    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, "Invalid request data", http.StatusBadRequest)
        return
    }

    points, err := calculatePoints(receipt)
    if err != nil {
        http.Error(w, "Error calculating points: "+err.Error(), http.StatusBadRequest)
        return
    }

    id := uuid.New().String()
    receipts[id] = points

    json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/receipts/")
    id = strings.TrimSuffix(id, "/points")

    points, exists := receipts[id]

    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(PointsResponse{Points: points})
}

func main() {
    http.HandleFunc("/receipts/process", processReceipt)
    http.HandleFunc("/receipts/", getPoints) 

    fmt.Printf("Starting server at port 8080\n")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}