package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simulacrumBank/internal/rest/client"
)

func InitRouters(e *gin.Engine) {
	//todo сделать мидлваре с каким нибудь логом
	initCommonRouters(e)
	client.InitClientRouters(e)
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
