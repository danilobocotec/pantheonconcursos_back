package model

type AsaasCustomerRequest struct {
	Name                 string `json:"name" binding:"required" example:"teste "`
	CPFOrCNPJ            string `json:"cpfCnpj" binding:"required" example:"616.236.260-48"`
	Email                string `json:"email" binding:"required,email" example:"guilherme.matossouza@gmail.com"`
	Phone                string `json:"phone" binding:"required" example:"(22) 2 2222-2222"`
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
