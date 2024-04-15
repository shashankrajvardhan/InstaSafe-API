package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var ist *time.Location
	ist = time.FixedZone("IST", int((5*time.Hour + 30*time.Minute).Seconds()))
	var ct string
	ct = time.Now().In(ist).Format(time.RFC3339)

	var jsonPayload Transaction
	jsonPayload = Transaction{
		Amount:    "100.25",
		Timestamp: ct,
	}

	payloadBytes, err := json.Marshal(jsonPayload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	var a string
	a = string(payloadBytes)
	fmt.Println("JSON Payload:", a)

	router()
}
