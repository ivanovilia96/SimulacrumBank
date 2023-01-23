package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	cash_account_actions "simulacrumBank/internal/actions/cash_account"
	client_actions "simulacrumBank/internal/actions/client"
	"simulacrumBank/internal/adapters/data_base"
	cash_account_handlers "simulacrumBank/internal/handlers/cash_account"
	client_handlers "simulacrumBank/internal/handlers/client"
)

func InitRouters(e *gin.Engine, dbActions data_base.DataBaseActions) {
	//todo сделать мидлваре с каким нибудь логом
	initCommonRouters(e)
	// то есть БД вызывается в Actions - они вызываются для handlers - а handlers в rest
	// получается gin такой же фуфил внешний как и бд, его нужно вынести и подумать над этим
	clientActions := client_actions.NewActions(dbActions)
	clientHandlers := client_handlers.NewHandlers(clientActions)

	client := e.Group("/client")
	client.POST("/add", clientHandlers.Add)
	client.DELETE("/delete/:mail", clientHandlers.Delete)
	client.POST("/update/:mail", clientHandlers.Update)
	client.GET("/get/:mail", clientHandlers.Get)

	cashAccountHandlers := cash_account_handlers.NewHandlers(cash_account_actions.NewActions(dbActions, clientActions))

	cashAccount := e.Group("/cash_account")
	cashAccount.POST("/create", cashAccountHandlers.Create)
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
