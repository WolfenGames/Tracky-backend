package controller

import (
	"Tracky/config"
	"Tracky/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTestDB() {
	_ = config.InitDb(":memory:")
}

func TestLogTrackable(t *testing.T) {
	setupTestDB()

	trackable := model.Trackable{
		Tool:      "VSCode",
		Metadata:  "path/to/file.txt",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(30 * time.Minute),
	}
	body, _ := json.Marshal(trackable)

	req := httptest.NewRequest(http.MethodPost, "/trackable", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(LogTrackable)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler message: %v", rr.Body)
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if message, ok := response["message"].(string); !ok || message != "Trackable logged" {
		t.Errorf("Unexpected message in response: %v", response)
	}
	if _, ok := response["id"].(float64); !ok {
		t.Errorf("expected an `id` field in response")
	}
}

func TestMissingFieldsTrackable(t *testing.T) {
	setupTestDB()

	trackable := `{"tool": "VSCode"}`
	req := httptest.NewRequest(http.MethodPost, "/trackable", bytes.NewReader([]byte(trackable)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogTrackable)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler message: %v", rr.Body)
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	if !bytes.Contains(rr.Body.Bytes(), []byte("Missing required fields")) {
		t.Errorf("Handler returned unexpected body: %v", rr.Body.String())
	}

}
