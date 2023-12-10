package routes

import (
	"coffee-delivery/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]models.User)

func SetupUserRouters(api *gin.Engine) {
	api.POST("/user", createUser)
	api.GET("/user", getAllUsers)
	api.GET("/user/:name", getUserName)

}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": db})
}

func getUserName(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func createUser(c *gin.Context) {
	// create user through body
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db[user.Email] = user
}
