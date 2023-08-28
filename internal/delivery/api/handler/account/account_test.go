package account

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	ginDriver "github.com/fms85/desafio-tecnico-go-stone/internal/driver/gin"
	accountUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/account"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestAccountHandlerGetAll(t *testing.T) {
	type dependencies struct {
		accountUsecase func() *accountUsecase.AccountUsecaseMock
	}
	tests := []struct {
		name         string
		dependencies dependencies
		want         string
		wantCode     int
	}{
		{
			name: "should_retrieve_accounts_successfully",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("GetAll", mock.Anything).Return([]*entity.Account{
						getAccountTest(1),
					}, nil)

					return usecase
				},
			},
			want:     `{"data":[{"id":1,"name":"Loren","cpf":"25462557035","balance":100,"createdAt":"0001-01-01T00:00:00Z"}]}`,
			wantCode: http.StatusOK,
		},
		{
			name: "should_return_an_error_when_usecase_get_all_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("GetAll", mock.Anything).Return(nil, errors.New(""))

					return usecase
				},
			},
			want:     `{"message":"internal server error"}`,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			handler := New(tt.dependencies.accountUsecase(), "")
			handler.InitRoutes(router)
			responseRecorder := httptest.NewRecorder()

			request, _ := http.NewRequest(
				"GET",
				"/accounts",
				nil,
			)

			router.ServeHTTP(responseRecorder, request)
			assert.Equal(t, tt.wantCode, responseRecorder.Code)

			if diff := cmp.Diff(responseRecorder.Body.String(), tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestAccountHandlerGetBalance(t *testing.T) {
	type dependencies struct {
		accountUsecase func() *accountUsecase.AccountUsecaseMock
	}
	tests := []struct {
		name         string
		dependencies dependencies
		accoundIdURI string
		want         string
		wantCode     int
	}{
		{
			name: "should_retrieve_account_balance_successfully",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1), nil)

					return usecase
				},
			},
			accoundIdURI: "1",
			want:         `{"balance":100}`,
			wantCode:     http.StatusOK,
		},
		{
			name: "should_return_an_error_when_usecase_get_all_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New(""))

					return usecase
				},
			},
			accoundIdURI: "1",
			want:         `{"message":"internal server error"}`,
			wantCode:     http.StatusInternalServerError,
		},
		{
			name: "should_return_an_error_validation_when_usecase_get_all_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(nil, &common.ValidationError{Msg: common.NOT_FOUND_ERROR})

					return usecase
				},
			},
			accoundIdURI: "1",
			want:         `{"message":"not found"}`,
			wantCode:     http.StatusBadRequest,
		},
		{
			name: "should_return_an_error_validation_when_should_bind_uri_get_all_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					return &accountUsecase.AccountUsecaseMock{}
				},
			},
			accoundIdURI: "x",
			want:         `{"message":"Key: 'GetBalanceAccountUri.AccountID' Error:Field validation for 'AccountID' failed on the 'numeric' tag"}`,
			wantCode:     http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			handler := New(tt.dependencies.accountUsecase(), "")
			handler.InitRoutes(router)
			responseRecorder := httptest.NewRecorder()

			request, _ := http.NewRequest(
				"GET",
				fmt.Sprintf("/accounts/%s/balance", tt.accoundIdURI),
				nil,
			)

			router.ServeHTTP(responseRecorder, request)
			assert.Equal(t, tt.wantCode, responseRecorder.Code)

			if diff := cmp.Diff(responseRecorder.Body.String(), tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestAccountHandlerCreateAccount(t *testing.T) {
	type dependencies struct {
		accountUsecase func() *accountUsecase.AccountUsecaseMock
	}
	tests := []struct {
		name         string
		dependencies dependencies
		body         []byte
		want         string
		wantCode     int
	}{
		{
			name: "should_retrieve_create_account_successfully",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Create", mock.Anything, mock.Anything).Return(getAccountTest(1), nil)

					return usecase
				},
			},
			body:     []byte(`{"name": "Loren","cpf": "25462557035","secret": "123456"}`),
			want:     `{"data":{"id":1,"name":"Loren","cpf":"25462557035","balance":100,"createdAt":"0001-01-01T00:00:00Z"}}`,
			wantCode: http.StatusCreated,
		},
		{
			name: "should_return_an_error_when_usecase_create_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New(""))

					return usecase
				},
			},
			body:     []byte(`{"name": "Loren","cpf": "25462557035","secret": "123456"}`),
			want:     `{"message":"internal server error"}`,
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "should_return_an_error_validation_when_usecase_create_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Create", mock.Anything, mock.Anything).Return(nil, &common.ValidationError{Msg: common.FOUND_ERROR})

					return usecase
				},
			},
			body:     []byte(`{"name": "Loren","cpf": "25462557035","secret": "123456"}`),
			want:     `{"message":"already exists"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "should_return_an_error_validation_when_should_bind_uri_create_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					return &accountUsecase.AccountUsecaseMock{}
				},
			},
			body:     []byte(`{"name": "Loren","cpf": "xxxxxxxxxx","secret": "123456"}`),
			want:     `{"message":"Key: 'AccountInput.CPF' Error:Field validation for 'CPF' failed on the 'cpf' tag"}`,
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := ginDriver.Setup()
			handler := New(tt.dependencies.accountUsecase(), "")
			handler.InitRoutes(router)
			responseRecorder := httptest.NewRecorder()

			request, _ := http.NewRequest(
				"POST",
				"/accounts",
				bytes.NewReader(tt.body),
			)

			router.ServeHTTP(responseRecorder, request)
			assert.Equal(t, tt.wantCode, responseRecorder.Code)

			if diff := cmp.Diff(responseRecorder.Body.String(), tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestAccountHandlerAuthAccount(t *testing.T) {
	type dependencies struct {
		accountUsecase func() *accountUsecase.AccountUsecaseMock
	}
	tests := []struct {
		name         string
		dependencies dependencies
		body         []byte
		want         string
		wantCode     int
	}{
		{
			name: "should_retrieve_auth_account_successfully",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(getAccountTest(1), nil)

					return usecase
				},
			},
			body:     []byte(`{"cpf": "25462557035","secret": "123456"}`),
			want:     `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoxfQ._2CpbtQXHkxxFD3qlzHtmJw8rvgPhIlnP42IAaZ9Kzk"}`,
			wantCode: http.StatusCreated,
		},
		{
			name: "should_return_an_error_when_usecase_auth_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(nil, errors.New(""))

					return usecase
				},
			},
			body:     []byte(`{"cpf": "25462557035","secret": "123456"}`),
			want:     `{"message":"internal server error"}`,
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "should_return_an_error_validation_when_usecase_auth_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					usecase := &accountUsecase.AccountUsecaseMock{}
					usecase.On("Get", mock.Anything, mock.Anything).Return(nil, &common.ValidationError{Msg: common.NOT_FOUND_ERROR})

					return usecase
				},
			},
			body:     []byte(`{"cpf": "25462557035","secret": "123456"}`),
			want:     `{"message":"not found"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "should_return_an_error_validation_when_should_bind_uri_auth_retrieval",
			dependencies: dependencies{
				accountUsecase: func() *accountUsecase.AccountUsecaseMock {
					return &accountUsecase.AccountUsecaseMock{}
				},
			},
			body:     []byte(`{"cpf": "xxxxxxx","secret": "123456"}`),
			want:     `{"message":"Key: 'CredentialsInput.CPF' Error:Field validation for 'CPF' failed on the 'cpf' tag"}`,
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := ginDriver.Setup()
			handler := New(tt.dependencies.accountUsecase(), "")
			handler.InitRoutes(router)
			responseRecorder := httptest.NewRecorder()

			request, _ := http.NewRequest(
				"POST",
				"/login",
				bytes.NewReader(tt.body),
			)

			router.ServeHTTP(responseRecorder, request)
			assert.Equal(t, tt.wantCode, responseRecorder.Code)

			if diff := cmp.Diff(responseRecorder.Body.String(), tt.want); diff != "" {
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
