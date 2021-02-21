package main

import (
	"log"

	"github.com/PatrikOlin/lp-api/controllers"
	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/middlewares"
	"github.com/PatrikOlin/lp-api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		public := v1.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		user := v1.Group("/user").Use(middlewares.Authz())
		{
			user.GET("/profile", controllers.GetAllProfiles)
			user.GET("/profile/:id", controllers.GetProfileByID)
			user.GET("/pickup/:id", controllers.GetPickupsByUserID)
			user.PUT("/profile", controllers.UpdateProfile)
			user.PUT("/recycler", controllers.ToggleRecycler)
		}

		haul := v1.Group("/haul").Use(middlewares.Authz())
		{
			haul.GET("/", controllers.GetAllHauls)
			haul.GET("/:id", controllers.GetHaulByID)
		}

		pickup := v1.Group("/pickup").Use(middlewares.Authz())
		{
			pickup.GET("/", controllers.GetAllPickups)
			pickup.GET("/:id", controllers.GetPickupByID)
			pickup.POST("/", controllers.CreatePickup)
			pickup.PUT("/:id", controllers.UpdatePickupByID)
			pickup.DELETE("/:id", controllers.DeletePickupByID)
		}
	}

	return r
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	err := db.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	db.GlobalDB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Haul{},
		&models.Proposition{},
		&models.Pickup{},
	)
}

func main() {
	Init()

	r := setupRouter()
	r.Run(":8125")

}
