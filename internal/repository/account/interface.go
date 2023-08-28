package account

import (
	"context"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
)

type IAccountRepository interface {
	GetAll(ctx context.Context) ([]*entity.Account, error)
	Get(ctx context.Context, accountInput types.AccountInput) (*entity.Account, error)
	Save(ctx context.Context, account *entity.Account) error
}
