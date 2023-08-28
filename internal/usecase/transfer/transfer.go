package transfer

import (
	"context"
	"fmt"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	transferRepository "github.com/fms85/desafio-tecnico-go-stone/internal/repository/transfer"
	accountUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/account"
	"github.com/fms85/desafio-tecnico-go-stone/internal/util"
)

type transferUsecase struct {
	transferRepository transferRepository.ITransferRepository
	accountUsecase     accountUsecase.IAccountUsecase
}

func New(transferRepository transferRepository.ITransferRepository, accountUsecase accountUsecase.IAccountUsecase) ITransferUsecase {
	return &transferUsecase{
		transferRepository: transferRepository,
		accountUsecase:     accountUsecase,
	}
}

func (usecase *transferUsecase) GetAll(ctx context.Context, transferInput types.TransferInput) ([]*entity.Transfer, error) {
	transfers, err := usecase.transferRepository.Get(ctx, transferInput)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}

func (usecase *transferUsecase) Create(ctx context.Context, transferInput types.TransferInput) error {
	if transferInput.AccountOriginID == transferInput.AccountDestinationID {
		return &common.ValidationError{Msg: "origin and destination accounts are equal"}
	}

	accountOrigin, err := usecase.accountUsecase.Get(ctx, types.AccountInput{ID: transferInput.AccountOriginID})
	if err != nil {
		return fmt.Errorf("account origin %w", &common.ValidationError{Msg: common.NOT_FOUND_ERROR})
	}

	accountDestination, err := usecase.accountUsecase.Get(ctx, types.AccountInput{ID: transferInput.AccountDestinationID})
	if err != nil {
		return fmt.Errorf("account destination %w", &common.ValidationError{Msg: common.NOT_FOUND_ERROR})
	}

	transferAggregation := types.CreateTransferAggregation(transferInput, accountOrigin, accountDestination)
	if err := usecase.setAmount(transferAggregation); err != nil {
		return err
	}

	if err := usecase.transferRepository.Create(ctx, transferAggregation); err != nil {
		return err
	}

	return nil
}

func (usecase *transferUsecase) setAmount(transferAggregation *types.TransferAggregation) error {
	if transferAggregation.Transfer.Amount > transferAggregation.AccountOrigin.Balance {
		return &common.ValidationError{Msg: "insufficient funds"}
	}

	transferAggregation.AccountOrigin.Balance = util.Sub(
		transferAggregation.AccountOrigin.Balance,
		transferAggregation.Transfer.Amount,
	)

	transferAggregation.AccountDestination.Balance = util.Sum(
		transferAggregation.AccountDestination.Balance,
		transferAggregation.Transfer.Amount,
	)

	return nil
}
