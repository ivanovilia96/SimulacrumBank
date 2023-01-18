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

type HandlersList interface {
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
}
