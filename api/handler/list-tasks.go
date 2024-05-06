package handler

import (
	"api/gen"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ListTasks(c *gin.Context) {
	tasks := []gen.Task{}
	c.JSON(http.StatusOK, tasks)
}
