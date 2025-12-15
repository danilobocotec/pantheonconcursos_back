package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

var _ model.VadeMecumJurisprudencia
var _ model.CapaVadeMecumJurisprudencia

// GetVadeMecumJurisprudencia godoc
// @Summary      Listar jurisprudências
// @Tags         vade-mecum-jurisprudencia
// @Produce      json
// @Success      200 {array} model.VadeMecumJurisprudencia
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia [get]
func (h *Handlers) GetVadeMecumJurisprudencia(c *gin.Context) {
	items, err := h.jurisprudenciaService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetCapasVadeMecumJurisprudencia godoc
// @Summary      Listar capas de jurisprudência ou buscar por nome específico
// @Tags         vade-mecum-jurisprudencia
// @Produce      json
// @Param        nomecodigo query string false "Filtro por nomecodigo"
// @Success      200 {array} model.CapaVadeMecumJurisprudencia
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/capas [get]
func (h *Handlers) GetCapasVadeMecumJurisprudencia(c *gin.Context) {
	nome := strings.TrimSpace(c.Query("nomecodigo"))

	if nome != "" {
		item, err := h.capaJurisService.GetByNomeCodigo(nome)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "capa não encontrada"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, []model.CapaVadeMecumJurisprudencia{*item})
		return
	}

	items, err := h.capaJurisService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// CreateVadeMecumJurisprudencia godoc
// @Summary      Criar jurisprudência
// @Tags         vade-mecum-jurisprudencia
// @Accept       json
// @Produce      json
// @Param        request body model.CreateVadeMecumJurisprudenciaRequest true "Dados do registro"
// @Success      201 {object} model.VadeMecumJurisprudencia
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia [post]
func (h *Handlers) CreateVadeMecumJurisprudencia(c *gin.Context) {
	var req model.CreateVadeMecumJurisprudenciaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.jurisprudenciaService.Create(&req)
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

// CreateCapaVadeMecumJurisprudencia godoc
// @Summary      Criar capa para jurisprudência
// @Tags         vade-mecum-jurisprudencia
// @Accept       json
// @Produce      json
// @Param        request body model.CreateCapaVadeMecumJurisprudenciaRequest true "Dados da capa"
// @Success      201 {object} model.CapaVadeMecumJurisprudencia
// @Failure      400 {object} map[string]string
// @Failure      409 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/capas [post]
func (h *Handlers) CreateCapaVadeMecumJurisprudencia(c *gin.Context) {
	var req model.CreateCapaVadeMecumJurisprudenciaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.capaJurisService.Create(&req)
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

// GetVadeMecumJurisprudenciaByID godoc
// @Summary      Buscar jurisprudência por ID
// @Tags         vade-mecum-jurisprudencia
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} model.VadeMecumJurisprudencia
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/{id} [get]
func (h *Handlers) GetVadeMecumJurisprudenciaByID(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	item, err := h.jurisprudenciaService.GetByID(id)
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

// UpdateVadeMecumJurisprudencia godoc
// @Summary      Atualizar jurisprudência
// @Tags         vade-mecum-jurisprudencia
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Param        request body model.UpdateVadeMecumJurisprudenciaRequest true "Campos para atualização"
// @Success      200 {object} model.VadeMecumJurisprudencia
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/{id} [put]
func (h *Handlers) UpdateVadeMecumJurisprudencia(c *gin.Context) {
	id := c.Param("id")

	var req model.UpdateVadeMecumJurisprudenciaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.jurisprudenciaService.Update(id, &req)
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

// UpdateCapaVadeMecumJurisprudencia godoc
// @Summary      Atualizar capa de jurisprudência
// @Tags         vade-mecum-jurisprudencia
// @Accept       json
// @Produce      json
// @Param        nomecodigo path string true "Identificador da capa"
// @Param        request body model.UpdateCapaVadeMecumJurisprudenciaRequest true "Campos para atualização"
// @Success      200 {object} model.CapaVadeMecumJurisprudencia
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/capas/{nomecodigo} [put]
func (h *Handlers) UpdateCapaVadeMecumJurisprudencia(c *gin.Context) {
	nome := c.Param("nomecodigo")

	var req model.UpdateCapaVadeMecumJurisprudenciaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.capaJurisService.Update(nome, &req)
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

// DeleteVadeMecumJurisprudencia godoc
// @Summary      Remover jurisprudência
// @Tags         vade-mecum-jurisprudencia
// @Param        id path string true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/{id} [delete]
func (h *Handlers) DeleteVadeMecumJurisprudencia(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.jurisprudenciaService.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ImportVadeMecumJurisprudencia godoc
// @Summary      Importar jurisprudências via Excel
// @Description  Recebe um arquivo .xlsx com colunas específicas (vademecum_Jurisprudencia)
// @Tags         vade-mecum-jurisprudencia
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "Arquivo Excel (.xlsx)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /vade-mecum/jurisprudencia/import [post]
func (h *Handlers) ImportVadeMecumJurisprudencia(c *gin.Context) {
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

	count, err := h.jurisprudenciaService.ImportFromExcel(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"imported": count})
}
