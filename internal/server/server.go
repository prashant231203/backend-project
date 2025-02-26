package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prashant231203/backend-project/internal/handlers"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/predict", handlers.Predict_handler).Methods("POST")

	log.Println("Server starting on :5552")
	if err := http.ListenAndServe(":5552", router); err != nil {
		log.Fatal(err)
	}
}
