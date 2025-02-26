package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Predict_handler(w http.ResponseWriter, r *http.Request) {
	// Log incoming request
	log.Printf("Received request with Content-Type: %s", r.Header.Get("Content-Type"))

	// Parse the input data
	var features []float64
	if err := json.NewDecoder(r.Body).Decode(&features); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, fmt.Sprintf("Invalid input format: %v", err), http.StatusBadRequest)
		return
	}

	// Log the received features
	log.Printf("Received features: %v", features)

	// TODO: Add your model prediction logic here
	// For now, let's just return a dummy response
	response := map[string]interface{}{
		"status":     "success",
		"prediction": 0.0,
		"features":   features,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
