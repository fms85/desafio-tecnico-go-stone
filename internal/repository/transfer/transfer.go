package transfer

import (
	"context"
	"fmt"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	"gorm.io/gorm"
)

type transferRepository struct {
	read  *gorm.DB
	write *gorm.DB
}

func New(connections map[string]*gorm.DB) ITransferRepository {
	return &transferRepository{
		write: connections["wr"],
		read:  connections["rd"],
	}
}

func (repo *transferRepository) Get(ctx context.Context, transferInput types.TransferInput) ([]*entity.Transfer, error) {
	var transfers []*entity.Transfer

	if err := repo.read.WithContext(ctx).Where(transferInput).Find(&transfers).Error; err != nil {
		return nil, fmt.Errorf("error to get all transfers: %w", err)
	}

	return transfers, nil
}

func (repo *transferRepository) Create(ctx context.Context, transferAggregation *types.TransferAggregation) error {
	return repo.write.Transaction(func(tx *gorm.DB) error {

		if err := tx.WithContext(ctx).Create(&transferAggregation.Transfer).Error; err != nil {
			return fmt.Errorf("error to create transfer: %w", err)
		}

		tx.WithContext(ctx).Model(&entity.Account{}).
			Where("id = ?", transferAggregation.AccountOrigin.ID).
			Updates(map[string]interface{}{
				"balance": transferAggregation.AccountOrigin.Balance,
			})

		if tx.Error != nil {
			return fmt.Errorf("error to update account origin: %w", tx.Error)
		}

		tx.WithContext(ctx).Model(&entity.Account{}).
			Where("id = ?", transferAggregation.AccountDestination.ID).
			Updates(map[string]interface{}{
				"balance": transferAggregation.AccountDestination.Balance,
			})

		if tx.Error != nil {
			return fmt.Errorf("error to update account destination: %w", tx.Error)
		}

		return nil
	})
}
