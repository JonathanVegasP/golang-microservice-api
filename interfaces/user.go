package interfaces

import (
	"flutter-store-api/domain/services"
	"flutter-store-api/infrastructure/auth"
	"flutter-store-api/infrastructure/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUser(service services.IUserController) *user {
	return &user{
		service: service,
	}
}

type user struct {
	service services.IUserController
}

func (u *user) Login(c *gin.Context) {
	var login dtos.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	user := u.service.GetUserByLogin(&login)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "E-mail/Password is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (u *user) CreateUser(c *gin.Context) {
	var user dtos.CreateUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	if success := u.service.CreateUser(&user); success != nil {
		c.JSON(http.StatusCreated, success)
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"message": "This email is already in use",
		})
	}
}

func (u *user) DeleteUser(c *gin.Context) {
	id := uint64(auth.GetClaims(c.Request)["sub"].(float64))

	if result := u.service.DeleteUser(&id); result {
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted successfully",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}
}

func (u *user) GetUser(c *gin.Context) {
	id := uint64(auth.GetClaims(c.Request)["sub"].(float64))
	user := u.service.GetUser(&id)
	if user != nil {
		fmt.Println(user.Email)
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}

}

func (u *user) UpdateUser(c *gin.Context) {
	var user dtos.UpdateUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	id := uint64(auth.GetClaims(c.Request)["sub"].(float64))

	if success := u.service.UpdateUser(&user, &id); success {
		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"message": "This e-mail is already in use",
		})
	}

}
