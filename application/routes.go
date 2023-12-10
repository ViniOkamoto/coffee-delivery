package application

import (
	"coffee-delivery/main/routes"

	"github.com/gin-gonic/gin"
)

func loadRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	routes.SetupUserRouters(router)

	return router
}
