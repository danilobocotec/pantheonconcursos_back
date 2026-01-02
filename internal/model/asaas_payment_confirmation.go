package model

type AsaasCreditCard struct {
	Number     string `json:"number" binding:"required"`
	ExpiryMonth string `json:"expiryMonth" binding:"required"`
	CCV        string `json:"ccv" binding:"required"`
	ExpiryYear string `json:"expiryYear" binding:"required"`
	HolderName string `json:"holderName" binding:"required"`
}

type AsaasCreditCardHolderInfo struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	CPFOrCNPJ     string `json:"cpfCnpj" binding:"required"`
	PostalCode    string `json:"postalCode" binding:"required"`
	AddressNumber string `json:"addressNumber" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
}

type AsaasPaymentConfirmationRequest struct {
	CreditCard           AsaasCreditCard           `json:"creditCard" binding:"required"`
	CreditCardHolderInfo AsaasCreditCardHolderInfo `json:"creditCardHolderInfo" binding:"required"`
}
