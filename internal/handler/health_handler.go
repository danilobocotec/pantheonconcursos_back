package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "API is running",
	})
}
