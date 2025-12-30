package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
)

// CreateVadeMecum godoc
// @Summary      Criar item de Vade Mecum
// @Tags         vade-mecum
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumRequest true "Dados do Vade Mecum"
// @Success      201 {object} model.VadeMecum
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum [post]
func (h *Handlers) CreateVadeMecum(c *gin.Context) {
	var req model.CreateVadeMecumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.vadeMecumService.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetVadeMecum godoc
// @Summary      Listar itens de Vade Mecum
// @Tags         vade-mecum
// @Produce      json
// @Success      200 {array} model.VadeMecum
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum [get]
func (h *Handlers) GetVadeMecum(c *gin.Context) {
	items, err := h.vadeMecumService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetVadeMecumByID godoc
// @Summary      Obter Vade Mecum por ID
// @Tags         vade-mecum
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecum
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/{id} [get]
func (h *Handlers) GetVadeMecumByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.vadeMecumService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateVadeMecum godoc
// @Summary      Atualizar Vade Mecum
// @Tags         vade-mecum
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumRequest true "Dados para atualização"
// @Success      200 {object} model.VadeMecum
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/{id} [put]
func (h *Handlers) UpdateVadeMecum(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateVadeMecumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.vadeMecumService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteVadeMecum godoc
// @Summary      Remover Vade Mecum
// @Tags         vade-mecum
// @Param        id path string true "ID"
// @Success      204 {object} nil
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/{id} [delete]
func (h *Handlers) DeleteVadeMecum(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.vadeMecumService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func normalizeCategory(param string) string {
	return strings.ToLower(param)
}

// CreateVadeMecumByCategory godoc
// @Summary      Criar item específico por categoria
// @Tags         vade-mecum
// @Accept       json
// @Produce      json
// @Param        category path string true "Categoria"
// @Param        request body model.CreateVadeMecumRequest true "Dados"
// @Success      201 {object} model.VadeMecum
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/category/{category} [post]
func (h *Handlers) CreateVadeMecumByCategory(c *gin.Context) {
	var req model.CreateVadeMecumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Category = normalizeCategory(c.Param("category"))

	item, err := h.vadeMecumService.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetVadeMecumByCategory godoc
// @Summary      Listar itens por categoria
// @Tags         vade-mecum
// @Produce      json
// @Param        category path string true "Categoria"
// @Success      200 {array} model.VadeMecum
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/category/{category} [get]
func (h *Handlers) GetVadeMecumByCategory(c *gin.Context) {
	category := normalizeCategory(c.Param("category"))
	items, err := h.vadeMecumService.GetByCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// UpdateVadeMecumByCategory godoc
// @Summary      Atualizar item dentro de uma categoria
// @Tags         vade-mecum
// @Accept       json
// @Produce      json
// @Param        category path string true "Categoria"
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumRequest true "Dados"
// @Success      200 {object} model.VadeMecum
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/category/{category}/{id} [put]
func (h *Handlers) UpdateVadeMecumByCategory(c *gin.Context) {
	category := normalizeCategory(c.Param("category"))
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.vadeMecumService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if item.Category != category {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in this category"})
		return
	}

	var req model.UpdateVadeMecumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Category = category

	updated, err := h.vadeMecumService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteVadeMecumByCategory godoc
// @Summary      Remover item de uma categoria
// @Tags         vade-mecum
// @Param        category path string true "Categoria"
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/category/{category}/{id} [delete]
func (h *Handlers) DeleteVadeMecumByCategory(c *gin.Context) {
	category := normalizeCategory(c.Param("category"))
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.vadeMecumService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if item.Category != category {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in this category"})
		return
	}

	if err := h.vadeMecumService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
