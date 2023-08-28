package transfer

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fms85/desafio-tecnico-go-stone/internal/delivery/api/middleware"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	transferUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/transfer"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestTransferandlerGetAll(t *testing.T) {
	type dependencies struct {
		transferUsecase func() *transferUsecase.TransferUsecaseMock
	}
	tests := []struct {
		name         string
		dependencies dependencies
		want         string
		wantCode     int
	}{
		{
			name: "should_retrieve_transfers_successfully",
			dependencies: dependencies{
				transferUsecase: func() *transferUsecase.TransferUsecaseMock {
					usecase := &transferUsecase.TransferUsecaseMock{}
					usecase.On("GetAll", mock.Anything, mock.Anything).Return([]*entity.Transfer{
						getTransferTest(1),
					}, nil)

					return usecase
				},
			},
			want:     `{"data":[{"id":1,"account_origin_id":1,"account_destination_id":2,"amount":10,"createdAt":"0001-01-01T00:00:00Z"}]}`,
			wantCode: http.StatusOK,
		},
		{
			name: "should_return_an_error_when_usecase_get_all_retrieval",
			dependencies: dependencies{
				transferUsecase: func() *transferUsecase.TransferUsecaseMock {
					usecase := &transferUsecase.TransferUsecaseMock{}
					usecase.On("GetAll", mock.Anything, mock.Anything).Return(nil, errors.New(""))

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
			authorized := router.Group("/")
			authorized.Use(middleware.Auth("2aa5b62a718429b0645dc1be1bcac023821181859a181408b59c77d7c07d5349"))

			handler := New(tt.dependencies.transferUsecase())
			handler.InitRoutes(authorized)
			responseRecorder := httptest.NewRecorder()

			request, _ := http.NewRequest(
				"GET",
				"/transfers",
				nil,
			)

			request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjozfQ.Wm6uJKmN9CO7f8224bZuICKojMkzvkXbr-EQlB13bz0")

			router.ServeHTTP(responseRecorder, request)
			assert.Equal(t, tt.wantCode, responseRecorder.Code)

			if diff := cmp.Diff(responseRecorder.Body.String(), tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestTransferandlerCreate(t *testing.T) {
	type dependencies struct {
		transferUsecase func() *transferUsecase.TransferUsecaseMock
	}
	tests := []struct {
		name         string
		dependencies dependencies
		body         []byte
		want         string
		wantCode     int
	}{
		{
			name: "should_retrieve_create_transfer_successfully",
			dependencies: dependencies{
				transferUsecase: func() *transferUsecase.TransferUsecaseMock {
					usecase := &transferUsecase.TransferUsecaseMock{}
					usecase.On("Create", mock.Anything, mock.Anything).Return(nil)

					return usecase
				},
			},
			body:     []byte(`{"account_destination_id": 2,"amount": 1}`),
			want:     `null`,
			wantCode: http.StatusCreated,
		},
		{
			name: "should_return_an_error_when_usecase_create_retrieval",
			dependencies: dependencies{
				transferUsecase: func() *transferUsecase.TransferUsecaseMock {
					usecase := &transferUsecase.TransferUsecaseMock{}
					usecase.On("Create", mock.Anything, mock.Anything).Return(errors.New(""))

					return usecase
				},
			},
			body:     []byte(`{"account_destination_id": 2,"amount": 1}`),
			want:     `{"message":"internal server error"}`,
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "should_return_an_error_validation_when_usecase_create_retrieval",
			dependencies: dependencies{
				transferUsecase: func() *transferUsecase.TransferUsecaseMock {
					usecase := &transferUsecase.TransferUsecaseMock{}
					usecase.On("Create", mock.Anything, mock.Anything).Return(&common.ValidationError{Msg: common.FOUND_ERROR})

					return usecase
				},
			},
			body:     []byte(`{"account_destination_id": 2,"amount": 1}`),
			want:     `{"message":"already exists"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "should_return_an_error_validation_when_should_bind_uri_create_retrieval",
			dependencies: dependencies{
				transferUsecase: func() *transferUsecase.TransferUsecaseMock {
					return &transferUsecase.TransferUsecaseMock{}
				},
			},
			body:     []byte(`{"amount": 1}`),
			want:     `{"message":"Key: 'TransferInput.AccountDestinationID' Error:Field validation for 'AccountDestinationID' failed on the 'required' tag"}`,
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			authorized := router.Group("/")
			authorized.Use(middleware.Auth("2aa5b62a718429b0645dc1be1bcac023821181859a181408b59c77d7c07d5349"))

			handler := New(tt.dependencies.transferUsecase())
			handler.InitRoutes(authorized)
			responseRecorder := httptest.NewRecorder()

			request, _ := http.NewRequest(
				"POST",
				"/transfers",
				bytes.NewReader(tt.body),
			)

			request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjozfQ.Wm6uJKmN9CO7f8224bZuICKojMkzvkXbr-EQlB13bz0")

			router.ServeHTTP(responseRecorder, request)
			assert.Equal(t, tt.wantCode, responseRecorder.Code)

			if diff := cmp.Diff(responseRecorder.Body.String(), tt.want); diff != "" {
				t.Error(diff)
			}
		})
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
