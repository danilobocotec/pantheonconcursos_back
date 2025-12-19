package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
)

// GetEstatutos godoc
// @Summary      Listar estatutos
// @Tags         vade-mecum-estatutos
// @Produce      json
// @Success      200 {array} model.VadeMecumEstatuto
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/estatutos [get]
func (h *Handlers) GetEstatutos(c *gin.Context) {
	items, err := h.estatutoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetEstatutoGrupoServico godoc
// @Summary      Listar grupos de estatutos por nomecodigo (grupo servico)
// @Tags         vade-mecum-estatutos
// @Produce      json
// @Success      200 {array} model.VadeMecumEstatutoGrupoServico
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/estatutos/gruposervico [get]
func (h *Handlers) GetEstatutoGrupoServico(c *gin.Context) {
	items, err := h.estatutoService.GrupoServico()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetEstatutoByID godoc
// @Summary      Obter estatuto por ID
// @Tags         vade-mecum-estatutos
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecumEstatuto
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/estatutos/{id} [get]
func (h *Handlers) GetEstatutoByID(c *gin.Context) {
	id, err := uuid.Parse(strings.TrimSpace(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	item, err := h.estatutoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateEstatuto godoc
// @Summary      Criar estatuto
// @Tags         vade-mecum-estatutos
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumEstatutoRequest true "Dados do estatuto"
// @Success      201 {object} model.VadeMecumEstatuto
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/estatutos [post]
func (h *Handlers) CreateEstatuto(c *gin.Context) {
	var req model.CreateVadeMecumEstatutoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.estatutoService.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateEstatuto godoc
// @Summary      Atualizar estatuto
// @Tags         vade-mecum-estatutos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumEstatutoRequest true "Campos para atualização"
// @Success      200 {object} model.VadeMecumEstatuto
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/estatutos/{id} [put]
func (h *Handlers) UpdateEstatuto(c *gin.Context) {
	id, err := uuid.Parse(strings.TrimSpace(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req model.UpdateVadeMecumEstatutoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.estatutoService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteEstatuto godoc
// @Summary      Remover estatuto
// @Tags         vade-mecum-estatutos
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/estatutos/{id} [delete]
func (h *Handlers) DeleteEstatuto(c *gin.Context) {
	id, err := uuid.Parse(strings.TrimSpace(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.estatutoService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
