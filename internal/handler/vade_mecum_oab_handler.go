package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

var _ model.VadeMecumOAB

// GetVadeMecumOAB godoc
// @Summary      Listar conteúdo OAB
// @Tags         vade-mecum-oab
// @Produce      json
// @Success      200 {array} model.VadeMecumOAB
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab [get]
func (h *Handlers) GetVadeMecumOAB(c *gin.Context) {
	items, err := h.oabService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// CreateCapaVadeMecumOAB godoc
// @Summary      Criar capa para vade-mécum OAB
// @Tags         vade-mecum-oab
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCapaVadeMecumOABRequest true "Dados da capa"
// @Success      201 {object} model.CapaVadeMecumOAB
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/capas [post]
func (h *Handlers) CreateCapaVadeMecumOAB(c *gin.Context) {
	var req model.CreateCapaVadeMecumOABRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.capaOABService.Create(&req)
	if err != nil {
		lower := strings.ToLower(err.Error())
		if strings.Contains(lower, "já existe") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetCapasVadeMecumOAB godoc
// @Summary      Listar capas OAB ou buscar por nome específico
// @Tags         vade-mecum-oab
// @Produce      json
// @Param        nomecodigo query string false "Filtro por nomecodigo"
// @Success      200 {array} model.CapaVadeMecumOAB
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/capas [get]
func (h *Handlers) GetCapasVadeMecumOAB(c *gin.Context) {
	nome := strings.TrimSpace(c.Query("nomecodigo"))

	if nome != "" {
		item, err := h.capaOABService.GetByNomeCodigo(nome)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "capa não encontrada"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, []model.CapaVadeMecumOAB{*item})
		return
	}

	items, err := h.capaOABService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// UpdateCapaVadeMecumOAB godoc
// @Summary      Atualizar capa OAB
// @Tags         vade-mecum-oab
// @Accept       json
// @Produce      json
// @Param        id path string true "Identificador da capa"
// @Param        request body model.UpdateCapaVadeMecumOABRequest true "Campos para atualização"
// @Success      200 {object} model.CapaVadeMecumOAB
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/capas/{id} [put]
func (h *Handlers) UpdateCapaVadeMecumOAB(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req model.UpdateCapaVadeMecumOABRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.capaOABService.Update(id, &req)
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

// CreateVadeMecumOAB godoc
// @Summary      Criar registro OAB
// @Tags         vade-mecum-oab
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumOABRequest true "Dados do registro"
// @Success      201 {object} model.VadeMecumOAB
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab [post]
func (h *Handlers) CreateVadeMecumOAB(c *gin.Context) {
	var req model.CreateVadeMecumOABRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.oabService.Create(&req)
	if err != nil {
		lower := strings.ToLower(err.Error())
		if strings.Contains(lower, "já existe") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetVadeMecumOABByID godoc
// @Summary      Obter registro OAB por ID
// @Tags         vade-mecum-oab
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecumOAB
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/{id} [get]
func (h *Handlers) GetVadeMecumOABByID(c *gin.Context) {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	item, err := h.oabService.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateVadeMecumOAB godoc
// @Summary      Atualizar registro OAB
// @Tags         vade-mecum-oab
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumOABRequest true "Campos para atualização"
// @Success      200 {object} model.VadeMecumOAB
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/{id} [put]
func (h *Handlers) UpdateVadeMecumOAB(c *gin.Context) {
	id := c.Param("id")

	var req model.UpdateVadeMecumOABRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.oabService.Update(id, &req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteVadeMecumOAB godoc
// @Summary      Remover registro OAB
// @Tags         vade-mecum-oab
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/{id} [delete]
func (h *Handlers) DeleteVadeMecumOAB(c *gin.Context) {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.oabService.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ImportVadeMecumOAB godoc
// @Summary      Importar conteúdo OAB via Excel
// @Description  Recebe um arquivo .xlsx para importar registros do Vade Mecum OAB
// @Tags         vade-mecum-oab
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "Arquivo Excel (.xlsx)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/oab/import [post]
func (h *Handlers) ImportVadeMecumOAB(c *gin.Context) {
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

	count, err := h.oabService.ImportFromExcel(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"imported": count})
}
