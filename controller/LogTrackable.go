package controller

import (
	"Tracky/config"
	"Tracky/model"
	"encoding/json"
	"io"
	"net/http"
)

func LogTrackable(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to get body: "+err.Error(), http.StatusBadRequest)
		return
	}
	var trackable model.Trackable
	if err = json.Unmarshal(body, &trackable); err != nil {
		http.Error(w, "Error parsing body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if trackable.Tool == "" || trackable.StartTime.IsZero() {
		http.Error(w, "Missing required fields: tool, start_time", http.StatusBadRequest)
		return
	}

	id, err := model.LogTrackable(db, trackable)
	if err != nil {
		http.Error(w, "Failed to log activity: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Trackable logged",
		"id":      id,
	})
}
