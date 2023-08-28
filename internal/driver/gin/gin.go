package gin

import (
	"github.com/fms85/desafio-tecnico-go-stone/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Setup() *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("cpf", util.CpfValidator); err != nil {
			panic(err)
		}
	}

	return router
}
