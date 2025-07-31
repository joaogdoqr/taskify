package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaogdoqr/taskfy/services"
)

type Application struct {
	Router      *chi.Mux
	TaskService services.TaskService
}
