package model

type AsaasCreditCard struct {
	Number     string `json:"number" binding:"required"`
	ExpiryMonth string `json:"expiryMonth" binding:"required"`
	CCV        string `json:"ccv" binding:"required"`
	ExpiryYear string `json:"expiryYear" binding:"required"`
	HolderName string `json:"holderName" binding:"required"`
}

type AsaasPaymentConfirmationRequest struct {
	CreditCard AsaasCreditCard `json:"creditCard" binding:"required"`
}
