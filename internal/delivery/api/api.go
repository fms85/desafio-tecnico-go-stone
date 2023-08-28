package api

import (
	"fmt"
	"net/http"
	"time"

	accountHandler "github.com/fms85/desafio-tecnico-go-stone/internal/delivery/api/handler/account"
	transferHandler "github.com/fms85/desafio-tecnico-go-stone/internal/delivery/api/handler/transfer"
	middleware "github.com/fms85/desafio-tecnico-go-stone/internal/delivery/api/middleware"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	accountRepository "github.com/fms85/desafio-tecnico-go-stone/internal/repository/account"
	transferRepository "github.com/fms85/desafio-tecnico-go-stone/internal/repository/transfer"
	accountUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/account"
	transferUsecase "github.com/fms85/desafio-tecnico-go-stone/internal/usecase/transfer"
	"github.com/fms85/desafio-tecnico-go-stone/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init(app common.App) {
	router := Setup()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Env.HTTP_ADDR),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	accountRepository := accountRepository.New(app.DB)
	transferRepository := transferRepository.New(app.DB)

	accountUsecase := accountUsecase.New(accountRepository)
	transferUsecase := transferUsecase.New(transferRepository, accountUsecase)

	accountHandler := accountHandler.New(accountUsecase, app.Env.JWT_SECRET)
	accountHandler.InitRoutes(router)

	authorized := router.Group("/")
	authorized.Use(middleware.Auth(app.Env.JWT_SECRET))
	{
		transferHandler := transferHandler.New(transferUsecase)
		transferHandler.InitRoutes(authorized)
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func Setup() *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("cpf", util.CpfValidator); err != nil {
			panic(err)
		}
	}

	return router
}
