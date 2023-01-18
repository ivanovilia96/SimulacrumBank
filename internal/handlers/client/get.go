package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) Get(c *gin.Context) {
	mail := c.Param("mail")

	r, err := h.Actions.Get(mail)
	if err != nil {
		c.JSON(666, err.Error())
		return
	}

	c.JSON(http.StatusOK, r)
}
