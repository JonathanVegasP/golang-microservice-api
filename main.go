package main

import (
	"flutter-store-api/application"
	"flutter-store-api/infrastructure/auth"
	"flutter-store-api/infrastructure/persistente"
	"flutter-store-api/interfaces"
	"flutter-store-api/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	auth.Init()

	db := persistente.Init()

	repo := persistente.NewUserRepository(db)

	service := application.NewUserService(repo)

	api := interfaces.NewUser(service)

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	v1 := r.Group("/v1")

	{
		v1.POST("/user", api.CreateUser)

		v1.POST("/login", api.Login)

		authenticated := v1.Group("/")
		{
			authenticated.Use(middlewares.AuthMiddleware())

			authenticated.GET("/user", api.GetUser)

			authenticated.PATCH("/user", api.UpdateUser)

			authenticated.DELETE("/user", api.DeleteUser)
		}
	}

	r.Run()
}
