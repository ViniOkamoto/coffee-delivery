package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Create(*gin.Context) {
	fmt.Println("User Created")
}

func (u *UserHandler) Update(*gin.Context) {
	fmt.Println("User Updated")
}

func (u *UserHandler) Delete(*gin.Context) {
	fmt.Println("User Deleted")
}

func (u *UserHandler) GetAll(*gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"users": db})

}

func (u *UserHandler) GetByID(*gin.Context) {
	fmt.Println("User Get By ID")
}

func CreateUserRoutes(api *gin.Engine) {
	handler := &UserHandler{}
	api.POST("/user", handler.Create)
	api.GET("/user", handler.GetAll)
	api.GET("/user/:id", handler.GetByID)

}

// func getAllUsers(c *gin.Context) {
// }

// func getUserName(c *gin.Context) {
// 	user := c.Params.ByName("name")
// 	value, ok := db[user]
// 	if ok {
// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 	}
// }

// func createUser(c *gin.Context) {
// 	// create user through body
// 	var user models.User
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	db[user.Email] = user
// }
