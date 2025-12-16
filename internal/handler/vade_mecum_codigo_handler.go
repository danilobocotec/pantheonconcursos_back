package handler

import (
	"errors"
	"net/http"
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
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

// CreateCapaVadeMecumCodigo godoc
// @Summary      Criar capa de vade-mécum código
// @Tags         vade-mecum-codigos
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCapaVadeMecumCodigoRequest true "Dados da capa"
// @Success      201 {object} model.CapaVadeMecumCodigo
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos/capas [post]
func (h *Handlers) CreateCapaVadeMecumCodigo(c *gin.Context) {
	var req model.CreateCapaVadeMecumCodigoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.capaCodigoService.Create(&req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "já existe") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetCapasVadeMecumCodigo godoc
// @Summary      Listar capas de códigos ou buscar por nome específico
// @Tags         vade-mecum-codigos
// @Produce      json
// @Param        nomecodigo query string false "Filtro por nomecodigo"
// @Success      200 {array} model.CapaVadeMecumCodigo
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos/capas [get]
func (h *Handlers) GetCapasVadeMecumCodigo(c *gin.Context) {
	nome := strings.TrimSpace(c.Query("nomecodigo"))

	if nome != "" {
		item, err := h.capaCodigoService.GetByNomeCodigo(nome)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "capa não encontrada"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, []model.CapaVadeMecumCodigo{*item})
		return
	}

	items, err := h.capaCodigoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// UpdateCapaVadeMecumCodigo godoc
// @Summary      Atualizar capa de vade-mécum código
// @Tags         vade-mecum-codigos
// @Accept       json
// @Produce      json
// @Param        id path string true "Identificador da capa"
// @Param        request body model.UpdateCapaVadeMecumCodigoRequest true "Campos para atualização"
// @Success      200 {object} model.CapaVadeMecumCodigo
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos/capas/{id} [put]
func (h *Handlers) UpdateCapaVadeMecumCodigo(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req model.UpdateCapaVadeMecumCodigoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.capaCodigoService.Update(id, &req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "capa não encontrada"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// GetCodigosGrouped godoc
// @Summary      Listar codigos agrupados por nomecodigo
// @Tags         vade-mecum-codigos
// @Produce      json
// @Param        priority query []string false "Prioridade dos grupos, pode ser separado por vírgula ou repetido"
// @Success      200 {array} model.VadeMecumCodigoGroup
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/codigos/grouped [get]
func (h *Handlers) GetCodigosGrouped(c *gin.Context) {
	priorities := extractPriorityOrder(c)

	groups, err := h.codigoService.GetGroupedByNomeCodigo(priorities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
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

<<<<<<< HEAD
func extractPriorityOrder(c *gin.Context) []string {
	values := c.QueryArray("priority")
	if len(values) == 0 {
		if raw := c.Query("priorities"); raw != "" {
			values = append(values, raw)
		}
	}

	if len(values) == 0 {
		return nil
	}

	normalized := make([]string, 0, len(values))
	seen := make(map[string]struct{})

	for _, raw := range values {
		parts := strings.Split(raw, ",")
		for _, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed == "" {
				continue
			}

			key := strings.ToLower(trimmed)
			if _, exists := seen[key]; exists {
				continue
			}

			normalized = append(normalized, trimmed)
			seen[key] = struct{}{}
		}
	}

	if len(normalized) == 0 {
		return nil
	}

	return normalized
}

=======
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
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
