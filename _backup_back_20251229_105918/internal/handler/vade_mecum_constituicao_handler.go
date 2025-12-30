package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
)

// GetConstituicoes godoc
// @Summary      Listar constituições
// @Tags         vade-mecum-constituicao
// @Produce      json
// @Success      200 {array} model.VadeMecumConstituicao
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/constituicao [get]
func (h *Handlers) GetConstituicoes(c *gin.Context) {
	items, err := h.constituicaoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetConstituicaoGrupoServico godoc
// @Summary      Listar grupos de constituição por titulo (grupo servico)
// @Tags         vade-mecum-constituicao
// @Produce      json
// @Success      200 {array} model.VadeMecumConstituicaoGrupoServico
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/constituicao/gruposervico [get]
func (h *Handlers) GetConstituicaoGrupoServico(c *gin.Context) {
	items, err := h.constituicaoService.GrupoServico()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetConstituicaoByID godoc
// @Summary      Obter constituição por ID
// @Tags         vade-mecum-constituicao
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecumConstituicao
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/constituicao/{id} [get]
func (h *Handlers) GetConstituicaoByID(c *gin.Context) {
	id, err := uuid.Parse(strings.TrimSpace(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	item, err := h.constituicaoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateConstituicao godoc
// @Summary      Criar constituição
// @Tags         vade-mecum-constituicao
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumConstituicaoRequest true "Dados"
// @Success      201 {object} model.VadeMecumConstituicao
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/constituicao [post]
func (h *Handlers) CreateConstituicao(c *gin.Context) {
	var req model.CreateVadeMecumConstituicaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.constituicaoService.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateConstituicao godoc
// @Summary      Atualizar constituição
// @Tags         vade-mecum-constituicao
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumConstituicaoRequest true "Campos para atualização"
// @Success      200 {object} model.VadeMecumConstituicao
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/constituicao/{id} [put]
func (h *Handlers) UpdateConstituicao(c *gin.Context) {
	id, err := uuid.Parse(strings.TrimSpace(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req model.UpdateVadeMecumConstituicaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.constituicaoService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteConstituicao godoc
// @Summary      Remover constituição
// @Tags         vade-mecum-constituicao
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/constituicao/{id} [delete]
func (h *Handlers) DeleteConstituicao(c *gin.Context) {
	id, err := uuid.Parse(strings.TrimSpace(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.constituicaoService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ImportConstituicao godoc
// @Summary      Importar constituição via Excel
// @Description  Importa registros utilizando um arquivo Excel (.xlsx) com cabeçalho padrão
// @Tags         vade-mecum-constituicao
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "Arquivo Excel (.xlsx)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/constituicao/import [post]
func (h *Handlers) ImportConstituicao(c *gin.Context) {
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

	count, err := h.constituicaoService.ImportConstituicao(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"imported": count,
	})
}
