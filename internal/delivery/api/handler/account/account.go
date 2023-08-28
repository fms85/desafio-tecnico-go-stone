package account

import (
	"errors"
	"log"
	"net/http"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	accountUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/account"
	"github.com/fms85/desafio-tecnico-go-stone/internal/util"
	"github.com/gin-gonic/gin"
)

var validationError *common.ValidationError

type AccountHandler struct {
	accountUsecase accountUsecase.IAccountUsecase
	secret         string
}

func New(accountUsecase accountUsecase.IAccountUsecase, secret string) *AccountHandler {
	return &AccountHandler{
		accountUsecase: accountUsecase,
		secret:         secret,
	}
}

func (handler *AccountHandler) InitRoutes(router *gin.Engine) {
	router.GET("accounts", handler.getAll)
	router.GET("accounts/:account_id/balance", handler.getBalance)
	router.POST("accounts", handler.createBalance)
	router.POST("login", handler.authAccount)
}

func (handler *AccountHandler) getAll(ctx *gin.Context) {
	accounts, err := handler.accountUsecase.GetAll(ctx.Request.Context())
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": accounts})
}

func (handler *AccountHandler) getBalance(ctx *gin.Context) {
	var uri types.GetBalanceAccountUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	account, err := handler.accountUsecase.Get(ctx.Request.Context(), types.AccountInput{ID: util.StringToUint(uri.AccountID)})
	if err != nil {
		if errors.As(err, &validationError) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": account.Balance})
}

func (handler *AccountHandler) createBalance(ctx *gin.Context) {
	var accountInput types.AccountInput
	if err := ctx.ShouldBindJSON(&accountInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	account, err := handler.accountUsecase.Create(ctx.Request.Context(), accountInput)
	if err != nil {
		if errors.As(err, &validationError) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": account})
}

func (handler *AccountHandler) authAccount(ctx *gin.Context) {
	var credentialsInput *types.CredentialsInput
	if err := ctx.ShouldBindJSON(&credentialsInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	account, err := handler.accountUsecase.Get(
		ctx.Request.Context(),
		types.AccountInput{CPF: credentialsInput.CPF, Secret: util.GenerateHash(credentialsInput.Secret)},
	)
	if err != nil {
		if errors.As(err, &validationError) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR})

		return
	}

	token := util.GenerateJwtToken("account_id", account.ID, handler.secret)

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}
