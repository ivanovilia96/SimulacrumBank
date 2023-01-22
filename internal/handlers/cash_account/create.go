package cash_account

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handlers) Create(c *gin.Context) {
	var d CreateCashAccountData
	err := c.BindJSON(&d)
	if err != nil {
		c.JSON(666, err.Error())
		return
	}

	id, err := h.Actions.Create(d.Mail, d.Currency)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, "client with this mail not found "+d.Mail)
			return
		}

		c.JSON(555, err.Error())
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("id of created cashAccount is :%d ", id))
}
