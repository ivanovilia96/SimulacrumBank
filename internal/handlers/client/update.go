package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update - updates client with "mail" which we pass in pas arg, to the request body value
func (h Handlers) Update(c *gin.Context) {
	mail := c.Param("mail")
	var dataForChange Client
	err := c.BindJSON(&dataForChange)
	if err != nil {
		println("BindJSON err", err.Error())
		c.JSON(http.StatusBadRequest, `err`)
		return
	}

	err = h.Actions.Update(mail, dataForChange.Mail, dataForChange.Fio, dataForChange.Age)
	if err != nil {
		c.JSON(666, err.Error())
		return
	}

	c.JSON(http.StatusOK, "updated")
}
