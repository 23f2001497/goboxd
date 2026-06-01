package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/thesouldev/goboxd/internal/api"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/run", api.RunHandler)

	log.Println("goboxd running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}