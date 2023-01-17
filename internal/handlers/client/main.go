package client

//todo эту часть я считаю "бизнес логикой"

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
	println("alling Add")
	var client Client
	err := c.BindJSON(&client)
	if err != nil {
		println("BindJSON err", err.Error())
		c.JSON(666, `err`)
		return
	}

	err = h.Actions.Add(client.Mail, client.Fio, client.Age)
	if err != nil {
		println("h.Actions.Add err", err.Error())
		c.JSON(666, `err`)
		return
	}

	c.JSON(200, "vse horosho "+client.Mail)
}

type HandlersList interface {
	Add(c *gin.Context)
}
