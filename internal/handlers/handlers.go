package handlers

import (
	"log/slog"
	"net/http"
)

/*
// Services below will be added in folder services later, just for example
type ExampleService struct {}
func (e *ExampleService) ExampleFunc() {}

type ExampleServices struct {
	*ExampleService
}

// We will be able to do this
var i = NewHandlers(logger, &ExampleServices{ExampleService: &ExampleService{}})
*/

// Struct with all interfaces needed for handlers
type Services interface {
	//HealthService // just for example
}

/*
struct for all hendlers in project, it consists of each group of handlers
(e.g user: GetByID, GetAll...), it is created to provide simple assembly in main.go
*/
type Handlers struct {
	logger *slog.Logger
	health *Health
}

func NewHandlers(logger *slog.Logger, s Services) *Handlers {
	return &Handlers{
		logger: logger,
		health: NewHealth(logger /*, s*/), // if services needed, wi will add s as a param
	}
}

// func to register each handler
func (h *Handlers) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.health.Health)
}
