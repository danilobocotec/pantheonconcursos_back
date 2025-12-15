package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
)

// CreateCodigo godoc
// @Summary      Criar item na categoria codigos
// @Tags         vade-mecum-codigos
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumCodigoRequest true "Dados do codigo"
// @Success      201 {object} model.VadeMecumCodigo
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos [post]
func (h *Handlers) CreateCodigo(c *gin.Context) {
	var req model.CreateVadeMecumCodigoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.codigoService.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetCodigos godoc
// @Summary      Listar codigos
// @Tags         vade-mecum-codigos
// @Produce      json
// @Success      200 {array} model.VadeMecumCodigo
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos [get]
func (h *Handlers) GetCodigos(c *gin.Context) {
	items, err := h.codigoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetCodigoByID godoc
// @Summary      Obter codigo por ID
// @Tags         vade-mecum-codigos
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecumCodigo
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/codigos/{id} [get]
func (h *Handlers) GetCodigoByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.codigoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateCodigo godoc
// @Summary      Atualizar codigo
// @Tags         vade-mecum-codigos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumCodigoRequest true "Dados"
// @Success      200 {object} model.VadeMecumCodigo
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/codigos/{id} [put]
func (h *Handlers) UpdateCodigo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateVadeMecumCodigoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.codigoService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteCodigo godoc
// @Summary      Remover codigo
// @Tags         vade-mecum-codigos
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos/{id} [delete]
func (h *Handlers) DeleteCodigo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.codigoService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ImportCodigos godoc
// @Summary      Importar codigos via Excel
// @Description  Importa registros utilizando um arquivo Excel (.xlsx) com cabeçalho padrão
// @Tags         vade-mecum-codigos
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "Arquivo Excel (.xlsx)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos/import [post]
func (h *Handlers) ImportCodigos(c *gin.Context) {
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

	count, err := h.codigoService.ImportFromExcel(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"imported": count,
	})
}
