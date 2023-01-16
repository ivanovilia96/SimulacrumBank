package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simulacrumBank/internal/adapters/data_base"
	"simulacrumBank/internal/rest/client"
)

func InitRouters(e *gin.Engine, dbActions data_base.DataBaseActions) {
	//todo сделать мидлваре с каким нибудь логом
	initCommonRouters(e)
	client.InitClientRouters(e, dbActions)
}

func initCommonRouters(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			ServiceName,
			Status string
		}{
			ServiceName: "simulacrumBank",
			Status:      "ok",
		})
	})
}
