package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

// CreateLei godoc
// @Summary      Criar lei
// @Tags         vade-mecum-leis
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumLeiRequest true "Dados da lei"
// @Success      201 {object} model.VadeMecumLei
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/leis [post]
func (h *Handlers) CreateLei(c *gin.Context) {
	var req model.CreateVadeMecumLeiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.leisService.Create(&req)
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

// GetLeis godoc
// @Summary      Listar leis
// @Tags         vade-mecum-leis
// @Produce      json
// @Success      200 {array} model.VadeMecumLei
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/leis [get]
func (h *Handlers) GetLeis(c *gin.Context) {
	items, err := h.leisService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetLeiByID godoc
// @Summary      Obter lei por ID
// @Tags         vade-mecum-leis
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecumLei
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /vade-mecum/leis/{id} [get]
func (h *Handlers) GetLeiByID(c *gin.Context) {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	item, err := h.leisService.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lei não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateLei godoc
// @Summary      Atualizar lei
// @Tags         vade-mecum-leis
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumLeiRequest true "Campos para atualização"
// @Success      200 {object} model.VadeMecumLei
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/leis/{id} [put]
func (h *Handlers) UpdateLei(c *gin.Context) {
	id := c.Param("id")

	var req model.UpdateVadeMecumLeiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.leisService.Update(id, &req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lei não encontrada"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteLei godoc
// @Summary      Remover lei
// @Tags         vade-mecum-leis
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/leis/{id} [delete]
func (h *Handlers) DeleteLei(c *gin.Context) {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.leisService.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lei não encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ImportLeis godoc
// @Summary      Importar leis via Excel
// @Description  Recebe um arquivo .xlsx e realiza o upsert das leis no banco
// @Tags         vade-mecum-leis
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "Arquivo Excel (.xlsx)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/leis/import [post]
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
