package controller

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health := HealthResponse{Status: "ok"}
	json.NewEncoder(w).Encode(&health)
}
