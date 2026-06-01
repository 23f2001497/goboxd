package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/thesouldev/goboxd/internal/executor"
	"github.com/thesouldev/goboxd/internal/models"
	"github.com/thesouldev/goboxd/internal/workspace"
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		return
	}

	if req.Language == "" || req.Source == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "missing language or source"})
		return
	}

	supported := map[string]bool{
		"python": true,
		"py3":    true,
	}
	if !supported[req.Language] {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "unsupported language"})
		return
	}

	dir, err := workspace.Create()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "workspace error"})
		return
	}
	defer workspace.Cleanup(dir)

	filePath := filepath.Join(dir, "solution.py")
	err = os.WriteFile(filePath, []byte(req.Source), 0644)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "file write error"})
		return
	}

	log.Printf("Written file to path: %s", filePath)

	res, err := executor.Run(filePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "execution error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
