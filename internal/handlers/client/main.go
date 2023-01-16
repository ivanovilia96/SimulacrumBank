package client

//todo эту часть я считаю "бизнес логикой"

import (
	"github.com/gin-gonic/gin"
	"io"
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
	var client Client
	err := c.BindJSON(&client)
	if err != nil {
		println("BindJSON err", err.Error())
		c.JSON(666, `err`)
	}
	err := h.Actions.Add()
	if err != nil {
		println("h.Actions.Add err", err.Error())
		c.JSON(666, `err`)
	}
}

type HandlersList interface {
	Add(c *gin.Context)
}
