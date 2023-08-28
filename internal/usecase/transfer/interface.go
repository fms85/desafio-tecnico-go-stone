package transfer

import (
	"context"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
)

type ITransferUsecase interface {
	GetAll(ctx context.Context, transferInput types.TransferInput) ([]*entity.Transfer, error)
	Create(ctx context.Context, transferInput types.TransferInput) error
}
