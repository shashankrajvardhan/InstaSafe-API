package main

import "net/http"

func router() {
	http.HandleFunc("/trans", handleTransaction)
	http.HandleFunc("/statistics", handleStatistics)
	http.ListenAndServe(":8080", nil)
}
