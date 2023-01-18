package client

import (
	"github.com/gin-gonic/gin"
	client_actions "simulacrumBank/internal/actions/client"
	"simulacrumBank/internal/adapters/data_base"
	client_handlers "simulacrumBank/internal/handlers/client"
)

func InitClientRouters(e *gin.Engine, dbActions data_base.DataBaseActions) {
	// то есть БД вызывается в Actions - они вызываются для handlers - а handlers в rest
	// получается gin такой же фуфил внешний как и бд, его нужно вынести и подумать над этим
	clientHandlers := client_handlers.NewHandlers(client_actions.NewActions(dbActions))

	client := e.Group("/client")
	client.POST("/add", clientHandlers.Add)
	client.POST("/delete", clientHandlers.Delete)
	client.POST("/update/:mail", clientHandlers.Update)
	client.GET("/get/:mail", clientHandlers.Get)
}
