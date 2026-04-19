package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	IP    string `json:"ip"`
	Count int    `json:"count"`
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {

	// 追加
	w.Header().Set("Access-Control-Allow-Origin", "*")

	results := []Result{
		{"192.168.1.1", 12},
		{"10.0.0.1", 7},
	}

	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/analyze", analyzeHandler)
	http.ListenAndServe(":8080", nil)
}
