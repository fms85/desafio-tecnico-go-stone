package account

import (
	"context"
	"fmt"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	accountRepository "github.com/fms85/desafio-tecnico-go-stone/internal/repository/account"
	"github.com/fms85/desafio-tecnico-go-stone/internal/util"
)

type accountUsecase struct {
	accountRepository accountRepository.IAccountRepository
}

func New(accountRepository accountRepository.IAccountRepository) IAccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
	}
}

func (usecase *accountUsecase) GetAll(ctx context.Context) ([]*entity.Account, error) {
	accounts, err := usecase.accountRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (usecase *accountUsecase) Get(ctx context.Context, accountInput types.AccountInput) (*entity.Account, error) {
	account, err := usecase.accountRepository.Get(ctx, accountInput)
	if err != nil {
		return nil, err
	}

	if account.ID == 0 {
		return nil, fmt.Errorf("account %w", &common.ValidationError{Msg: common.NOT_FOUND_ERROR})
	}

	return account, nil
}

func (usecase *accountUsecase) Create(ctx context.Context, accountInput types.AccountInput) (*entity.Account, error) {
	account, err := usecase.accountRepository.Get(ctx, types.AccountInput{CPF: accountInput.CPF})
	if err != nil {
		return nil, err
	}

	if account.ID > 0 {
		return nil, fmt.Errorf("account %w", &common.ValidationError{Msg: common.FOUND_ERROR})
	}

	account = &entity.Account{
		Name:    accountInput.Name,
		CPF:     accountInput.CPF,
		Secret:  util.GenerateHash(accountInput.Secret),
		Balance: entity.ENTITY_BALANCE_DEFAULT,
	}

	if err := usecase.accountRepository.Save(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}
