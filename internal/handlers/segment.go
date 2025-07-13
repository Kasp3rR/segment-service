package handlers

import (
	"encoding/json"
	"net/http"

	"segment-service/internal/models"

	"github.com/go-chi/chi/v5"
)

type CreateSegmentRequest struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	DistributionRatio float64 `json:"distribution_ratio"` // от 0 до 1
}

func CreateSegmentHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateSegmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	segment := models.Segment{
		Name:              req.Name,
		Description:       req.Description,
		DistributionRatio: req.DistributionRatio,
	}

	err := models.CreateSegment(segment)
	if err != nil {
		http.Error(w, "failed to create segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	err := models.DeleteSegment(name)
	if err != nil {
		http.Error(w, "failed to delete segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AssignSegmentRandomlyHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	var payload struct {
		Ratio float64 `json:"ratio"` // 0.0 - 1.0
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if payload.Ratio <= 0 || payload.Ratio > 1 {
		http.Error(w, "ratio must be between 0 and 1", http.StatusBadRequest)
		return
	}

	count, err := models.AssignSegmentRandomly(name, payload.Ratio)
	if err != nil {
		http.Error(w, "failed to assign segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"assigned": count})
}

func UpdateSegmentHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	var req struct {
		Description       string  `json:"description"`
		DistributionRatio float64 `json:"distribution_ratio"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if req.DistributionRatio < 0 || req.DistributionRatio > 1 {
		http.Error(w, "distribution_ratio must be between 0 and 1", http.StatusBadRequest)
		return
	}

	err := models.UpdateSegment(name, req.Description, req.DistributionRatio)
	if err != nil {
		http.Error(w, "failed to update segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetSegmentUsersHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	userIDs, err := models.GetSegmentUsers(name)
	if err != nil {
		http.Error(w, "failed to get segment users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"segment": name,
		"users":   userIDs,
		"count":   len(userIDs),
	})
}
