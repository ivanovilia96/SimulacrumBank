package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) Get(c *gin.Context) {
	var client Client

	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	r, err := h.Actions.Get(client.Mail)
	if err != nil {
		c.JSON(666, err.Error())
		return
	}

	c.JSON(http.StatusOK, r)
}
