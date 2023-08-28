package account

import (
	"context"
	"fmt"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	"gorm.io/gorm"
)

type accountRepository struct {
	read  *gorm.DB
	write *gorm.DB
}

func New(connections map[string]*gorm.DB) IAccountRepository {
	return &accountRepository{
		write: connections["wr"],
		read:  connections["rd"],
	}
}

func (repo *accountRepository) GetAll(ctx context.Context) ([]*entity.Account, error) {
	var accounts []*entity.Account

	if err := repo.read.WithContext(ctx).Find(&accounts).Error; err != nil {
		return nil, fmt.Errorf("error to get all account: %w", err)
	}

	return accounts, nil
}

func (repo *accountRepository) Get(ctx context.Context, accountInput types.AccountInput) (*entity.Account, error) {
	var account *entity.Account

	if err := repo.read.WithContext(ctx).Where(accountInput).Find(&account).Error; err != nil {
		return nil, fmt.Errorf("error to get account: %w", err)
	}

	return account, nil
}

func (repo *accountRepository) Save(ctx context.Context, account *entity.Account) error {
	if err := repo.write.WithContext(ctx).Save(&account).Error; err != nil {
		return fmt.Errorf("error to save account: %w", err)
	}

	return nil
}
