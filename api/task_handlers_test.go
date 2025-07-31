package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateTask(t *testing.T) {
	api := Application{}

	payload := map[string]any{
		"title":       "test",
		"description": "test",
		"priority":    8000,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload")
	}

	req := httptest.NewRequest("POST", "/api/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(api.handleCreateTask)
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Statuscode differs, expected %d, got %d", http.StatusCreated, rec.Code)
	}

	var resBody map[string]any
	err = json.Unmarshal(rec.Body.Bytes(), &body)
	if err != nil {
		t.Fatal("Failed to unmarshal response body")
	}

	if resBody["title"] != payload["title"] {
		t.Errorf("Title differs, expected %v, got %v", payload["title"], resBody["title"])
	}
}
