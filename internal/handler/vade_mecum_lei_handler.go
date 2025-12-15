package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) ImportLeis(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não enviado"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível abrir o arquivo"})
		return
	}
	defer src.Close()

	count, err := h.leisService.ImportFromExcel(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"imported": count})
}
