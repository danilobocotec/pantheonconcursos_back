package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/service"
)

// @Summary      Criar cliente Asaas
// @Description  Cria um cliente no Asaas e armazena o retorno localmente
// @Tags         asaas
// @Accept       json
// @Produce      json
// @Param        request  body      model.AsaasCustomerRequest  true  "Dados do cliente"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /asaas/customers [post]
func (h *Handlers) CreateAsaasCustomer(c *gin.Context) {
	var req model.AsaasCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, status, err := h.asaasService.CreateCustomer(c.Request.Context(), req)
	if err != nil {
		if apiErr, ok := err.(*service.AsaasAPIError); ok {
			c.JSON(apiErr.Status, gin.H{"error": apiErr.Body})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asaasID, ok := resp["id"].(string)
	if !ok || asaasID == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": "asaas response missing customer id"})
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer := model.AsaasCustomer{
		AsaasID:      asaasID,
		Name:         req.Name,
		CPFOrCNPJ:    req.CPFOrCNPJ,
		Email:        req.Email,
		Phone:        req.Phone,
		ResponseJSON: string(respJSON),
	}

	if err := h.asaasCustomerService.Create(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{
		"customer": resp,
		"stored":   customer,
	})
}

// @Summary      Criar pagamento Asaas (cartao)
// @Description  Cria um pagamento por cartao no Asaas
// @Tags         asaas
// @Accept       json
// @Produce      json
// @Param        request  body      model.AsaasPaymentRequest  true  "Dados do pagamento"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /asaas/payments [post]
func (h *Handlers) CreateAsaasPayment(c *gin.Context) {
	var req model.AsaasPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.asaasCustomerService.GetByAsaasID(req.Customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not registered"})
		return
	}

	req.BillingType = "CREDIT_CARD"

	reqJSON, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, status, err := h.asaasService.CreateCreditCardPayment(c.Request.Context(), req)
	if err != nil {
		if apiErr, ok := err.(*service.AsaasAPIError); ok {
			c.JSON(apiErr.Status, gin.H{"error": apiErr.Body})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	asaasID, ok := resp["id"].(string)
	if !ok || asaasID == "" {
		c.JSON(http.StatusBadGateway, gin.H{"error": "asaas response missing payment id"})
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payment := model.AsaasPayment{
		AsaasID:      asaasID,
		CustomerID:   req.Customer,
		Value:        req.Value,
		DueDate:      req.DueDate,
		BillingType:  req.BillingType,
		RequestJSON:  string(reqJSON),
		ResponseJSON: string(respJSON),
	}

	if err := h.asaasPaymentService.Create(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{
		"payment": resp,
		"stored":  payment,
	})
}

// @Summary      Confirmar pagamento Asaas (cartao)
// @Description  Confirma um pagamento de cartao no Asaas
// @Tags         asaas
// @Accept       json
// @Produce      json
// @Param        id       path      string                             true  "ID do pagamento Asaas"
// @Param        request  body      model.AsaasPaymentConfirmationRequest  true  "Dados de confirmacao"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /asaas/payments/{id}/confirm [post]
func (h *Handlers) ConfirmAsaasCreditCardPayment(c *gin.Context) {
	paymentID := c.Param("id")
	if paymentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment id is required"})
		return
	}

	var req model.AsaasPaymentConfirmationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.asaasPaymentService.GetByAsaasID(paymentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not registered"})
		return
	}

	resp, status, err := h.asaasService.ConfirmCreditCardPayment(c.Request.Context(), paymentID, req)
	if err != nil {
		if apiErr, ok := err.(*service.AsaasAPIError); ok {
			c.JSON(apiErr.Status, gin.H{"error": apiErr.Body})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.asaasPaymentService.UpdateConfirmation(paymentID, string(respJSON)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{
		"confirmation": resp,
	})
}

// @Summary      Criar webhook Asaas
// @Description  Cria um webhook no Asaas
// @Tags         asaas
// @Accept       json
// @Produce      json
// @Param        request  body      model.AsaasWebhookRequest  true  "Dados do webhook"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /asaas/webhooks [post]
func (h *Handlers) CreateAsaasWebhook(c *gin.Context) {
	var req model.AsaasWebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, status, err := h.asaasService.CreateWebhook(c.Request.Context(), req)
	if err != nil {
		if apiErr, ok := err.(*service.AsaasAPIError); ok {
			c.JSON(apiErr.Status, gin.H{"error": apiErr.Body})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{
		"webhook": resp,
	})
}
