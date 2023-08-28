package transfer

import (
	"errors"
	"log"
	"net/http"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/types"
	transferUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/transfer"
	"github.com/fms85/desafio-tecnico-go-stone/internal/util"
	"github.com/gin-gonic/gin"
)

var validationError *common.ValidationError

type TransferHandler struct {
	transferUsecase transferUsecase.ITransferUsecase
}

func New(transferUsecase transferUsecase.ITransferUsecase) *TransferHandler {
	return &TransferHandler{
		transferUsecase: transferUsecase,
	}
}

func (handler *TransferHandler) InitRoutes(router *gin.RouterGroup) {
	router.GET("transfers", handler.getAll)
	router.POST("transfers", handler.createTransfer)
}

func (handler *TransferHandler) getAll(ctx *gin.Context) {
	accountID := ctx.MustGet("account_id").(string)

	transfers, err := handler.transferUsecase.GetAll(ctx.Request.Context(), types.TransferInput{AccountOriginID: util.StringToUint(accountID)})
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transfers})
}

func (handler *TransferHandler) createTransfer(ctx *gin.Context) {
	var transferInput types.TransferInput
	if err := ctx.ShouldBindJSON(&transferInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	accountID := ctx.MustGet("account_id").(string)
	transferInput.AccountOriginID = util.StringToUint(accountID)

	err := handler.transferUsecase.Create(ctx.Request.Context(), transferInput)
	if err != nil {
		if errors.As(err, &validationError) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": common.INTERNAL_SERVER_ERROR})

		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
