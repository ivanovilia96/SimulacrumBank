package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) Delete(c *gin.Context) {
	var client Client
	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.Actions.Delete(client.Mail)
	if err != nil {
		c.JSON(666, err.Error())
		return
	}

	println("delete client with mail", client.Mail, "success")
	c.JSON(200, "email deleted:"+client.Mail)
}
