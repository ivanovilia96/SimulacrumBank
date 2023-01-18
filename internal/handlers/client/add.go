package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) Add(c *gin.Context) {
	println("alling Add")
	var client Client
	err := c.BindJSON(&client)
	if err != nil {
		println("BindJSON err", err.Error())
		c.JSON(http.StatusBadRequest, `err`)
		return
	}

	err = h.Actions.Add(client.Mail, client.Fio, client.Age)
	if err != nil {
		println("h.Actions.Add err", err.Error())
		c.JSON(666, `err`)
		return
	}

	c.JSON(201, "client created mail"+client.Mail)
}
