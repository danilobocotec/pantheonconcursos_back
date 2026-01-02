package model

type AsaasCustomerRequest struct {
	Name                 string `json:"name" binding:"required"`
	CPFOrCNPJ            string `json:"cpfCnpj" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Phone                string `json:"phone" binding:"required"`
	MobilePhone          string `json:"mobilePhone,omitempty"`
	ExternalReference    string `json:"externalReference,omitempty"`
	NotificationDisabled bool   `json:"notificationDisabled,omitempty"`
}

type AsaasPaymentRequest struct {
	Customer          string  `json:"customer" binding:"required"`
	BillingType       string  `json:"billingType,omitempty"`
	Value             float64 `json:"value" binding:"required,gt=0"`
	DueDate           string  `json:"dueDate" binding:"required"`
	Description       string  `json:"description,omitempty"`
	ExternalReference string  `json:"externalReference,omitempty"`
}

type AsaasWebhookRequest struct {
	Name        string   `json:"name" binding:"required"`
	URL         string   `json:"url" binding:"required,url"`
	Email       string   `json:"email,omitempty" binding:"omitempty,email"`
	Enabled     *bool    `json:"enabled" binding:"required"`
	Interrupted *bool    `json:"interrupted" binding:"required"`
	AuthToken   *string  `json:"authToken,omitempty"`
	SendType    string   `json:"sendType,omitempty"`
	Events      []string `json:"events" binding:"required,min=1,dive,required"`
}
