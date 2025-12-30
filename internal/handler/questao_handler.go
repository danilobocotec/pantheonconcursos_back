package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
)

// GetQuestoes godoc
// @Summary      Listar questoes
// @Tags         questoes
// @Produce      json
// @Param        disciplina query string false "Disciplina"
// @Param        assunto query string false "Assunto"
// @Param        banca query string false "Banca"
// @Param        orgao query string false "Orgao"
// @Param        cargo query string false "Cargo"
// @Param        concurso query string false "Concurso"
// @Param        area_conhecimento query string false "Area de conhecimento"
// @Success      200 {array} model.Questao
// @Failure      500 {object} map[string]string
// @Router       /questoes [get]
func (h *Handlers) GetQuestoes(c *gin.Context) {
	filters := buildQuestaoFilters(c)
	items, err := h.questaoService.GetAll(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetQuestaoFilters godoc
// @Summary      Listar filtros de questoes
// @Tags         questoes
// @Produce      json
// @Success      200 {object} model.QuestaoFiltersResponse
// @Failure      500 {object} map[string]string
// @Router       /questoes/filtros [get]
func (h *Handlers) GetQuestaoFilters(c *gin.Context) {
	items, err := h.questaoService.GetFilterOptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetQuestaoByID godoc
// @Summary      Obter questao por ID
// @Tags         questoes
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} model.Questao
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /questoes/{id} [get]
func (h *Handlers) GetQuestaoByID(c *gin.Context) {
	id, ok := parseQuestaoID(c)
	if !ok {
		return
	}

	item, err := h.questaoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateQuestao godoc
// @Summary      Criar questao
// @Tags         questoes
// @Accept       json
// @Produce      json
// @Param        request body model.CreateQuestaoRequest true "Dados da questao"
// @Success      201 {object} model.Questao
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /questoes [post]
func (h *Handlers) CreateQuestao(c *gin.Context) {
	var req model.CreateQuestaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.questaoService.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateQuestao godoc
// @Summary      Atualizar questao
// @Tags         questoes
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        request body model.UpdateQuestaoRequest true "Campos para atualizacao"
// @Success      200 {object} model.Questao
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /questoes/{id} [put]
func (h *Handlers) UpdateQuestao(c *gin.Context) {
	id, ok := parseQuestaoID(c)
	if !ok {
		return
	}

	var req model.UpdateQuestaoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.questaoService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteQuestao godoc
// @Summary      Remover questao
// @Tags         questoes
// @Param        id path int true "ID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /questoes/{id} [delete]
func (h *Handlers) DeleteQuestao(c *gin.Context) {
	id, ok := parseQuestaoID(c)
	if !ok {
		return
	}

	if err := h.questaoService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func parseQuestaoID(c *gin.Context) (int, bool) {
	idStr := strings.TrimSpace(c.Param("id"))
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return 0, false
	}
	return id, true
}

func buildQuestaoFilters(c *gin.Context) *model.QuestaoFilters {
	disciplina := strings.TrimSpace(c.Query("disciplina"))
	assunto := strings.TrimSpace(c.Query("assunto"))
	banca := strings.TrimSpace(c.Query("banca"))
	orgao := strings.TrimSpace(c.Query("orgao"))
	cargo := strings.TrimSpace(c.Query("cargo"))
	concurso := strings.TrimSpace(c.Query("concurso"))
	areaConhecimento := strings.TrimSpace(c.Query("area_conhecimento"))

	var filters model.QuestaoFilters
	hasAny := false

	if disciplina != "" {
		filters.Disciplina = &disciplina
		hasAny = true
	}
	if assunto != "" {
		filters.Assunto = &assunto
		hasAny = true
	}
	if banca != "" {
		filters.Banca = &banca
		hasAny = true
	}
	if orgao != "" {
		filters.Orgao = &orgao
		hasAny = true
	}
	if cargo != "" {
		filters.Cargo = &cargo
		hasAny = true
	}
	if concurso != "" {
		filters.Concurso = &concurso
		hasAny = true
	}
	if areaConhecimento != "" {
		filters.AreaConhecimento = &areaConhecimento
		hasAny = true
	}

	if !hasAny {
		return nil
	}

	return &filters
}
