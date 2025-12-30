package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
)

// GetPlans godoc
// @Summary      Listar planos com usuários vinculados
// @Description  Retorna todos os planos disponíveis e, para cada um, os usuários atrelados (quando houver)
// @Tags         plans
// @Produce      json
// @Success      200  {array}   model.Plan
// @Failure      500  {object}  map[string]string
// @Router       /plans [get]
func (h *Handlers) GetPlans(c *gin.Context) {
	var (
		plans []model.Plan
		err   error
	)

	if plans, err = h.planService.GetPlansWithUsers(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plans)
}
