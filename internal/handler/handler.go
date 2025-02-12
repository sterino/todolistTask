package handler

import (
	"github.com/gin-gonic/gin"
	"todoList/internal/handler/http"
	"todoList/internal/service/task"
	"todoList/pkg/server/router"
)

type Dependencies struct {
	TaskService *task.Service
}

type Configuration func(h *Handler) error

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
		taskHandler.Routes(r.Group("/tasks"))

		return
	}
}
