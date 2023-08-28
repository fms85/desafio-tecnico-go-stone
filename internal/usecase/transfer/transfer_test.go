package transfer

import (
	"context"
	"errors"
	"testing"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	transferRepository "github.com/fms85/desafio-tecnico-go-stone/internal/repository/transfer"
	accountUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/account"
	"github.com/google/go-cmp/cmp"
	mock "github.com/stretchr/testify/mock"
)

type dependencies struct {
	transferRepository func() *transferRepository.TransferRepositoryMock
	accountUsecase     func() *accountUsecase.AccountUsecaseMock
}

func TestTransferUsecaseGet(t *testing.T) {
	tests := []struct {
		name         string
		dependencies dependencies
		want         []*entity.Transfer
		wantErr      bool
	}{
		{
			name: "should_retrieve_transfers_successfully",
			dependencies: dependencies{
				transferRepository: func() *transferRepository.TransferRepositoryMock {
					repo := &transferRepository.TransferRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return([]*entity.Transfer{
						getTransferTest(1),
					}, nil)

					return repo
				},
			},
			want: []*entity.Transfer{
				getTransferTest(1),
			},
			wantErr: false,
		},
		{
			name: "should_return_an_error_when_repository_get_retrieval",
			dependencies: dependencies{
				transferRepository: func() *transferRepository.TransferRepositoryMock {
					repo := &transferRepository.TransferRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New(""))

					return repo
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := New(tt.dependencies.transferRepository(), &accountUsecase.AccountUsecaseMock{})

			got, err := usecase.GetAll(context.Background(), types.TransferInput{})
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestTransferUsecaseCreate(t *testing.T) {
	type params struct {
		transferInput types.TransferInput
	}

	tests := []struct {
		name         string
		dependencies dependencies
		params       params
		wantErr      bool
	}{
		{
			name: "should_retrieve_transfers_successfully",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1, 100), nil).Once()
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(2, 100), nil).Once()

					return usecase
				},

				transferRepository: func() *transferRepository.TransferRepositoryMock {
					repo := &transferRepository.TransferRepositoryMock{}
					repo.On("Create", mock.Anything, mock.Anything).Return(nil)

					return repo
				},
			},
			params: params{
				transferInput: getTransferInputTest(1, 2),
			},
			wantErr: false,
		},
		{
			name: "should_return_an_error_when_origin_and_destination_accounts_are_equal",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					return &accountUsecase.AccountUsecaseMock{}
				},

				transferRepository: func() *transferRepository.TransferRepositoryMock {
					return &transferRepository.TransferRepositoryMock{}
				},
			},
			params: params{
				transferInput: getTransferInputTest(1, 1),
			},
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_insufficient_funds",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1, 0), nil).Once()
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(2, 100), nil).Once()

					return usecase
				},

				transferRepository: func() *transferRepository.TransferRepositoryMock {
					return &transferRepository.TransferRepositoryMock{}
				},
			},
			params: params{
				transferInput: getTransferInputTest(1, 2),
			},
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_usecase_origin_get_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New("")).Once()

					return usecase
				},
				transferRepository: func() *transferRepository.TransferRepositoryMock {
					return &transferRepository.TransferRepositoryMock{}
				},
			},
			params: params{
				transferInput: getTransferInputTest(1, 2),
			},
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_usecase_destination_get_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1, 100), nil).Once()
					usecase.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New("")).Once()

					return usecase
				},
				transferRepository: func() *transferRepository.TransferRepositoryMock {
					return &transferRepository.TransferRepositoryMock{}
				},
			},
			params: params{
				transferInput: getTransferInputTest(1, 2),
			},
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_repository_get_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1, 100), nil).Once()
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(2, 100), nil).Once()

					return usecase
				},
				transferRepository: func() *transferRepository.TransferRepositoryMock {
					repo := &transferRepository.TransferRepositoryMock{}
					repo.On("Create", mock.Anything, mock.Anything).Return(errors.New(""))

					return repo
				},
			},
			params: params{
				transferInput: getTransferInputTest(1, 2),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := New(tt.dependencies.transferRepository(), tt.dependencies.accountUsecase())

			err := usecase.Create(context.Background(), tt.params.transferInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getAccountTest(id uint, balance float64) *entity.Account {
	return &entity.Account{
		ID:      id,
		Name:    "Loren",
		CPF:     "25462557035",
		Secret:  "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
		Balance: balance,
	}
}

func getTransferTest(id uint) *entity.Transfer {
	return &entity.Transfer{
		ID:                   id,
		AccountOriginID:      1,
		AccountDestinationID: 2,
		Amount:               10,
	}
}

func getTransferInputTest(accountOriginID uint, accountDestinationID uint) types.TransferInput {
	return types.TransferInput{
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               10,
	}
}
