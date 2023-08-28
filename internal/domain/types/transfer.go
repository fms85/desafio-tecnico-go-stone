package types

import "github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"

type TransferInput struct {
	AccountOriginID      uint
	AccountDestinationID uint    `json:"account_destination_id" binding:"required"`
	Amount               float64 `json:"amount" binding:"required"`
}

type TransferAggregation struct {
	AccountOrigin      *entity.Account
	AccountDestination *entity.Account
	Transfer           *entity.Transfer
}

func CreateTransferAggregation(TransferInput TransferInput, accountOrigin *entity.Account, accountDestination *entity.Account) *TransferAggregation {
	TransferAggregation := &TransferAggregation{}

	TransferAggregation.AccountOrigin = accountOrigin
	TransferAggregation.AccountDestination = accountDestination
	TransferAggregation.Transfer = &entity.Transfer{
		AccountOriginID:      TransferInput.AccountOriginID,
		AccountDestinationID: TransferInput.AccountDestinationID,
		Amount:               TransferInput.Amount,
	}

	return TransferAggregation
}
