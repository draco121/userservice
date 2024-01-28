package main

import (
	"github.com/draco121/common/database"
	"github.com/draco121/userservice/controllers"
	"github.com/draco121/userservice/core"
	"github.com/draco121/userservice/repository"
	"github.com/draco121/userservice/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func RunApp() {
	db := database.NewMongoDatabase(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DBNAME"))
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
