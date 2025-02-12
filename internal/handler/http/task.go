package http

import (
	"github.com/gin-gonic/gin"
	"todoList/internal/service/task"
)

type TaskHandler struct {
	taskService *task.Service
}

func NewTaskHandler(s *task.Service) *TaskHandler {
	return &TaskHandler{taskService: s}
}

func (h *TaskHandler) Routes(router *gin.RouterGroup) {

}
