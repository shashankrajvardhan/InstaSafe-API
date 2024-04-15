package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var transaction Transaction
	var err error
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	timestamp, err := time.Parse(time.RFC3339, transaction.Timestamp)
	if err != nil {
		http.Error(w, "Invalid timestamp format", http.StatusUnprocessableEntity)
		return
	}
	if timestamp.After(time.Now()) {
		http.Error(w, "Transaction date is in the future", http.StatusUnprocessableEntity)
		return
	}
	if time.Since(timestamp) > 60*time.Second {
		http.Error(w, "Transaction is older than 60 sec", http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Transaction successfully processed")
}
