package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/thesouldev/goboxd/internal/models"
	"github.com/thesouldev/goboxd/internal/workspace"
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	dir, err := workspace.Create()
	if err != nil {
		http.Error(w, "workspace error", 500)
		return
	}
	defer workspace.Cleanup(dir)

	filePath := filepath.Join(dir, "solution.py")
	err = os.WriteFile(filePath, []byte(req.Source), 0644)
	if err != nil {
		http.Error(w, "file write error", 500)
		return
	}

	log.Printf("Written file to path: %s", filePath)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "accepted",
	})
}
