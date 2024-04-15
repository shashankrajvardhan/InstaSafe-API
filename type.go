package main

type Transaction struct {
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type Statistics struct {
	Sum   string `json:"sum"`
	Avg   string `json:"avg"`
	Max   string `json:"max"`
	Min   string `json:"min"`
	Count string `json:"count"`
}