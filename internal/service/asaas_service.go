package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/thepantheon/api/internal/model"
)

type AsaasService struct {
	baseURL string
	token   string
	client  *http.Client
}

type AsaasAPIError struct {
	Status int
	Body   string
}

func (e *AsaasAPIError) Error() string {
	return fmt.Sprintf("asaas api error (%d): %s", e.Status, e.Body)
}

func NewAsaasService(baseURL, token string) *AsaasService {
	return &AsaasService{
		baseURL: strings.TrimRight(baseURL, "/"),
		token:   token,
		client: &http.Client{Timeout: 15 * time.Second},
	}
}

func (s *AsaasService) CreateCustomer(ctx context.Context, payload model.AsaasCustomerRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	status, err := s.doRequest(ctx, http.MethodPost, "/v3/customers", payload, &resp)
	return resp, status, err
}

func (s *AsaasService) CreatePayment(ctx context.Context, payload model.AsaasPaymentRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	status, err := s.doRequest(ctx, http.MethodPost, "/v3/payments", payload, &resp)
	return resp, status, err
}

func (s *AsaasService) CreateCreditCardPayment(ctx context.Context, payload model.AsaasPaymentRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	status, err := s.doRequest(ctx, http.MethodPost, "/v3/lean/payments", payload, &resp)
	return resp, status, err
}

func (s *AsaasService) ConfirmCreditCardPayment(ctx context.Context, paymentID string, payload model.AsaasPaymentConfirmationRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	path := fmt.Sprintf("/v3/payments/%s/payWithCreditCard", paymentID)
	status, err := s.doRequest(ctx, http.MethodPost, path, payload, &resp)
	return resp, status, err
}

func (s *AsaasService) CreateWebhook(ctx context.Context, payload model.AsaasWebhookRequest) (map[string]interface{}, int, error) {
	var resp map[string]interface{}
	status, err := s.doRequest(ctx, http.MethodPost, "/v3/webhooks", payload, &resp)
	return resp, status, err
}

func (s *AsaasService) doRequest(ctx context.Context, method, path string, payload interface{}, out interface{}) (int, error) {
	if s.baseURL == "" || s.token == "" {
		return http.StatusInternalServerError, fmt.Errorf("asaas configuration is missing")
	}

	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, s.baseURL+path, body)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("access_token", s.token)

	resp, err := s.client.Do(req)
	if err != nil {
		return http.StatusBadGateway, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadGateway, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return resp.StatusCode, &AsaasAPIError{
			Status: resp.StatusCode,
			Body:   strings.TrimSpace(string(respBody)),
		}
	}

	if out != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, out); err != nil {
			return http.StatusBadGateway, err
		}
	}

	return resp.StatusCode, nil
}
