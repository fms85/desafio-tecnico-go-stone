package account

import (
	"context"
	"errors"
	"testing"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	accountRepository "github.com/fms85/desafio-tecnico-go-stone/internal/repository/account"
	"github.com/google/go-cmp/cmp"
	mock "github.com/stretchr/testify/mock"
)

type dependencies struct {
	accountRepository func() *accountRepository.AccountRepositoryMock
}

func TestAccountUsecaseGetAll(t *testing.T) {
	tests := []struct {
		name         string
		dependencies dependencies
		want         []*entity.Account
		wantErr      bool
	}{
		{
			name: "should_retrieve_accounts_successfully",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("GetAll", mock.Anything).Return([]*entity.Account{
						getAccountTest(1),
					}, nil)

					return repo
				},
			},
			want: []*entity.Account{
				getAccountTest(1),
			},
			wantErr: false,
		},
		{
			name: "should_return_an_error_when_repository_get_all_retrieval",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("GetAll", mock.Anything).Return(nil, errors.New(""))

					return repo
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := New(tt.dependencies.accountRepository())

			got, err := usecase.GetAll(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestAccountUsecaseGet(t *testing.T) {
	tests := []struct {
		name         string
		dependencies dependencies
		want         *entity.Account
		wantErr      bool
	}{
		{
			name: "should_retrieve_an_account_successfully",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1), nil)

					return repo
				},
			},
			want:    getAccountTest(1),
			wantErr: false,
		},
		{
			name: "should_return_an_error_when_repository_get_retrieval",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New(""))

					return repo
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_account_is_not_found",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(&entity.Account{ID: 0}, nil)

					return repo
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := New(tt.dependencies.accountRepository())

			got, err := usecase.Get(context.Background(), types.AccountInput{ID: 1})
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestAccountUsecaseCreate(t *testing.T) {
	tests := []struct {
		name         string
		dependencies dependencies
		want         *entity.Account
		wantErr      bool
	}{
		{
			name: "should_create_an_account_successfully",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(&entity.Account{ID: 0}, nil)
					repo.On("Save", mock.Anything, mock.Anything).Return(nil)

					return repo
				},
			},

			want:    getAccountTest(0),
			wantErr: false,
		},
		{
			name: "should_return_an_error_when_repository_get_retrieval",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New(""))

					return repo
				},
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_account_already_exists",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(&entity.Account{ID: 1}, nil)

					return repo
				},
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "should_return_an_error_when_repository_save_retrieval",
			dependencies: dependencies{
				accountRepository: func() *accountRepository.AccountRepositoryMock {
					repo := &accountRepository.AccountRepositoryMock{}
					repo.On("Get", mock.Anything, mock.Anything).Return(&entity.Account{ID: 0}, nil)
					repo.On("Save", mock.Anything, mock.Anything).Return(errors.New(""))

					return repo
				},
			},

			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := New(tt.dependencies.accountRepository())

			got, err := usecase.Create(context.Background(), getAccountInputTest())
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func getAccountTest(id uint) *entity.Account {
	return &entity.Account{
		ID:      id,
		Name:    "Loren",
		CPF:     "25462557035",
		Secret:  "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
		Balance: 100,
	}
}

func getAccountInputTest() types.AccountInput {
	return types.AccountInput{
		Name:   "Loren",
		CPF:    "25462557035",
		Secret: "123456",
	}
}
