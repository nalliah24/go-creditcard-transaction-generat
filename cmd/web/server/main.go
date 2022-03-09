package main

import (
	"log"
	"net/http"

	th "github.com/nalliah24/go-creditcard-transaction-generator/cmd/web/handlers"
)

func homeHander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "API server is running. Use POST method with proper request to generate transactions. /api/transactions"}`))
	}
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHander)
	mux.HandleFunc("/api/transactions", th.TransHandler)

	log.Println("Starting server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
