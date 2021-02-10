package main

import (
	"log"

	"github.com/PatrikOlin/lp-api/controllers"
	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/middlewares"
	"github.com/PatrikOlin/lp-api/models"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/v1")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("signup", controllers.Signup)
		}

		user := api.Group("/user").Use(middlewares.Authz())
		{
			user.GET("/profile", controllers.Profile)
		}
	}

	return r
}

func main() {
	err := db.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	db.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()
	r.Run(":8125")

}
