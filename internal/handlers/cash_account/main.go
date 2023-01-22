package cash_account

import (
	"github.com/gin-gonic/gin"
	"simulacrumBank/internal/adapters/bisenes_logic"
)

func NewHandlers(a bisenes_logic.CashAccountActions) HandlersList {
	return Handlers{
		Actions: a,
	}
}

type HandlersList interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	AddMoney(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
