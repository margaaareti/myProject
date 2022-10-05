package authhttp

import (
	"github.com/gin-gonic/gin"
	"myProject/auth"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) SignUp(c *gin.Context) {

}
