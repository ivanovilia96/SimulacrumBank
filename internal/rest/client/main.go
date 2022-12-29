package client

import (
	"github.com/gin-gonic/gin"
	client_actions "simulacrumBank/internal/actions/client"
	client_handlers "simulacrumBank/internal/handlers/client"
)

func InitClientRouters(e *gin.Engine) {
	// то есть БД вызывается в Actions - они вызываются для handlers - а handlers в rest
	// получается gin такой же фуфил внешний как и бд, его нужно вынести и подумать над этим
	clientHandlers := client_handlers.NewHandlers(client_actions.NewActions())

	e.POST("/client/add", clientHandlers.Add)
}
