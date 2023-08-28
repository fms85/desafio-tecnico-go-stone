package types

type AccountInput struct {
	ID     uint
	Name   string `json:"name" binding:"required"`
	CPF    string `json:"cpf" binding:"required,cpf"`
	Secret string `json:"secret" binding:"required,min=6,max=12"`
}

type GetBalanceAccountUri struct {
	AccountID string `uri:"account_id" binding:"required,numeric"`
}

type CredentialsInput struct {
	CPF    string `json:"cpf" binding:"required,cpf"`
	Secret string `json:"secret" binding:"required,min=6,max=12"`
}
