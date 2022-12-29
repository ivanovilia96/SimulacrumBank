package client

import (
	"github.com/gin-gonic/gin"
	"simulacrumBank/internal/adapters/bisenes_logic"
)

func NewHandlers(a bisenes_logic.ClientActions) HandlersList {
	return Handlers{
		Actions: a,
	}
}

type Handlers struct {
	Actions bisenes_logic.ClientActions
}

func (h Handlers) Add(c *gin.Context) {
	h.Actions.Add()
}

type HandlersList interface {
	Add(c *gin.Context)
}
