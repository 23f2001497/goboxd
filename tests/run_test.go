package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/thesouldev/goboxd/internal/api"
)

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	api.HealthzHandler(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200 got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json got %s", contentType)
	}
}

func TestRunMissingLanguage(t *testing.T) {
	body := strings.NewReader(`{}`)
	req := httptest.NewRequest("POST", "/run", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	api.RunHandler(w, req)

	if w.Code != 400 {
		t.Errorf("expected 400 got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json got %s", contentType)
	}
}

func TestRunSuccess(t *testing.T) {
	body := strings.NewReader(`{"language": "python", "source": "print('hello')"}`)
	req := httptest.NewRequest("POST", "/run", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	api.RunHandler(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200 got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json got %s", contentType)
	}

	expected := `{"stdout":"hello\n"}`
	actual := strings.TrimSpace(w.Body.String())
	if actual != expected {
		t.Errorf("expected %s got %s", expected, actual)
	}
}
