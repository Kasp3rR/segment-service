package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"segment-service/internal/models"
)

type AssignSegmentRequest struct {
	SegmentName string `json:"segment_name"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := models.CreateUser(payload.UserID); err != nil {
		http.Error(w, "could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func AddUserToSegmentHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	var req AssignSegmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	segment, err := models.GetSegmentByName(req.SegmentName)
	if err != nil || segment == nil {
		http.Error(w, "segment not found", http.StatusNotFound)
		return
	}

	err = models.AddUserToSegment(userID, segment.ID)
	if err != nil {
		http.Error(w, "failed to add user to segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetUserSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	segments, err := models.GetUserSegments(userID)
	if err != nil {
		http.Error(w, "could not fetch user segments", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(segments)
}
