package routes

import (
	"github.com/draco121/userservice/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(controllers controllers.Controllers, router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.POST("/user", controllers.CreateTenant)
	v1.POST("/admin", controllers.CreateTenantAdmin)
	v1.GET("/user", controllers.GetUser)
	v1.PATCH("/user", controllers.UpdateUser)
	v1.DELETE("/user", controllers.DeleteUser)
}
