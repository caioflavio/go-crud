package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func main() {
	appPort := ":8080"
	log.Println("Servidor iniciado na porta: " + appPort)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		health := HealthResponse{Message: "hot reloading ok", StatusCode: 200}
		json.NewEncoder(w).Encode(health)
	})

	http.ListenAndServe(appPort, nil)
}
