package main

import (
	"log"
	"net/http"

	"github.com/thesouldev/goboxd/internal/api"
)

func main() {
	http.HandleFunc("/healthz", api.HealthzHandler)
	http.HandleFunc("/run", api.RunHandler)

	log.Println("goboxd running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}