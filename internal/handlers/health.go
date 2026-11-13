package handlers

import (
	"log/slog"
	"net/http"
)

// example
/*
type HealthService interface {
	// required methods
	// e.g. GetUserName(id int) (string, error)
	ExampleFunc()
}
*/

type Health struct {
	logger *slog.Logger
	//service HealthService // example
}

func NewHealth(logger *slog.Logger /*, s HealthService*/) *Health {
	return &Health{
		logger: logger,
		//service: s, // example
	}
}

func (h *Health) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
