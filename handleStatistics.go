package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var transactions []Transaction

func handleStatistics(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var (
		sum, max, min float64
		count         int
	)

	var ct time.Time
	ct = time.Now().UTC()
	for _, txn := range transactions {
		timestamp, err := time.Parse(time.RFC3339, txn.Timestamp)
		if err != nil {
			http.Error(w, "Error parsing timestamp", http.StatusInternalServerError)
			return
		}

		var diff time.Duration
		diff = ct.Sub(timestamp)
		if diff <= 60*time.Second {
			amount, err := strconv.ParseFloat(txn.Amount, 64)
			if err != nil {
				http.Error(w, "Error parsing amount", http.StatusInternalServerError)
				return
			}

			sum += amount

			if amount > max {
				max = amount
			}
			if count == 0 || amount < min {
				min = amount
			}
			count++
		}
	}

	var avg float64
	if count > 0 {
		avg = sum / float64(count)
	}

	var statistics Statistics
	statistics = Statistics{
		Sum:   fmt.Sprintf("%.2f", sum),
		Avg:   fmt.Sprintf("%.2f", avg),
		Max:   fmt.Sprintf("%.2f", max),
		Min:   fmt.Sprintf("%.2f", min),
		Count: strconv.Itoa(count),
	}

	response, err := json.Marshal(statistics)
	if err != nil {
		http.Error(w, "Error encoding statistics", http.StatusInternalServerError)
		return
	}

	var a http.Header
	a = w.Header()
	a.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
