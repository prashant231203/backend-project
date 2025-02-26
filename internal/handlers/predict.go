package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/prashant231203/backend-project/internal/models"
)

func PredictHandler(w http.ResponseWriter, r *http.Request) {
	var input []float64
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := models.GetPredictions(input)
	if err != nil {
		http.Error(w, "Prediction failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
