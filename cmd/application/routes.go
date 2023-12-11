package application

import (
	handler "coffee-delivery/main/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func loadRoutes() *gin.Engine {
	ConnectDatabase()

	router := gin.Default()
	router.Use(gin.Logger())

	handler.CreateUserRoutes(router)

	return router
}
