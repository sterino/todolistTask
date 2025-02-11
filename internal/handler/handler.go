package handler

import (
	"github.com/gin-gonic/gin"
	"todoList/pkg/server/router"
)

type Dependencies struct {
	TaskService *task.Service
}

type Configuration func(h *Handler) error

// Handler is an implementation of the Handler
type Handler struct {
	dependencies Dependencies

	HTTP *gin.Engine
}

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	h = &Handler{
		dependencies: d,
	}

	for _, cfg := range configs {

		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.HTTP = router.New()

		taskHandler := http.NewTaskHandler(h.dependencies.TaskService)

		r := h.HTTP.Group("/api")
		r.Group("/tasks", taskHandler)

		return
	}
}
