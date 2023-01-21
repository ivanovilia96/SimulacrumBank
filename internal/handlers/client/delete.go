package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simulacrumBank/internal/actions/client"
)

func (h Handlers) Delete(c *gin.Context) {
	mail := c.Param("mail")

	err := h.Actions.Delete(mail)
	if err != nil {
		if err == client.ErrorNoDeletedRows {
			c.JSON(http.StatusNotFound, err.Error()+"email is: "+mail)
			return
		}

		c.JSON(666, err.Error())
		return
	}

	println("delete client with mail", mail, "success")
	c.JSON(200, "email deleted:"+mail)
}
