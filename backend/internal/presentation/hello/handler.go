package hello

import (
	"net/http"

	"myapp/internal/usecase"
	"myapp/pkg/validator"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	GetHelloUsecase *usecase.GetHelloUsecase
}

func NewHandler(u *usecase.GetHelloUsecase) *Handler {
	return &Handler{
		GetHelloUsecase: u,
	}
}

func (h Handler) HelloWorld(ctx *gin.Context) {
	var param HelloWorldParams
	lang := ctx.DefaultQuery("lang", "ja")
	param.Lang = lang
	validate := validator.GetValidator()
	if err := validate.Struct(&param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	helloDTO, err := h.GetHelloUsecase.Exec(ctx.Request.Context(), param.Lang)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if helloDTO == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	response := HelloWorldResponse{
		Lang:    helloDTO.Lang,
		Message: helloDTO.Message,
	}

	ctx.JSON(http.StatusOK, response)
}
