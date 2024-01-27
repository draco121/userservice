package main

import (
	"userservice/controllers"
	"userservice/core"
	"userservice/repository"
	"userservice/routes"

	"github.com/draco121/common/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func RunApp() {
	db := database.NewMongoDatabaseDefaults()
	repo := repository.NewUserRepository(db)
	service := core.NewUserService(repo)
	controllers := controllers.NewControllers(service)
	router := gin.Default()
	routes.RegisterRoutes(controllers, router)
	err := router.Run()
	if err != nil {
		return
	}
}
func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	RunApp()
}
