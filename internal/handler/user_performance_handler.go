package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
)

// CreateUserPerformance godoc
// @Summary      Registrar desempenho do usuario
// @Tags         meu-desempenho
// @Accept       json
// @Produce      json
// @Param        request body model.CreateUserPerformanceRequest true "Desempenho"
// @Success      201 {object} model.UserPerformance
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meu-desempenho [post]
func (h *Handlers) CreateUserPerformance(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	var req model.CreateUserPerformanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.userPerformanceService.Create(userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetUserPerformance godoc
// @Summary      Listar desempenho do usuario
// @Tags         meu-desempenho
// @Produce      json
// @Param        data_inicio query string false "Data inicial (YYYY-MM-DD ou RFC3339)"
// @Param        data_fim query string false "Data final (YYYY-MM-DD ou RFC3339)"
// @Success      200 {array} model.UserPerformance
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meu-desempenho [get]
func (h *Handlers) GetUserPerformance(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	startDate, err := parsePerformanceDateParam(c.Query("data_inicio"), false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endDate, err := parsePerformanceDateParam(c.Query("data_fim"), true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := h.userPerformanceService.GetByUser(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetUserPerformanceSummary godoc
// @Summary      Resumo do desempenho do usuario
// @Tags         meu-desempenho
// @Produce      json
// @Param        data_inicio query string false "Data inicial (YYYY-MM-DD ou RFC3339)"
// @Param        data_fim query string false "Data final (YYYY-MM-DD ou RFC3339)"
// @Success      200 {object} model.UserPerformanceSummary
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /meu-desempenho/resumo [get]
func (h *Handlers) GetUserPerformanceSummary(c *gin.Context) {
	userID, ok := h.getUserIDFromRequest(c)
	if !ok {
		return
	}

	startDate, err := parsePerformanceDateParam(c.Query("data_inicio"), false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endDate, err := parsePerformanceDateParam(c.Query("data_fim"), true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	summary, err := h.userPerformanceService.GetSummary(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func parsePerformanceDateParam(value string, endOfDay bool) (*time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, nil
	}

	if len(value) == len("2006-01-02") {
		parsed, err := time.Parse("2006-01-02", value)
		if err != nil {
			return nil, errors.New("data invalida")
		}
		if endOfDay {
			parsed = parsed.Add(24*time.Hour - time.Nanosecond)
		}
		return &parsed, nil
	}

	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, errors.New("data invalida")
	}
	return &parsed, nil
}
