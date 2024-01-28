package controllers

import (
	"github.com/draco121/common/constants"
	"github.com/draco121/userservice/core"

	"github.com/draco121/common/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controllers struct {
	service core.IUserService
}

func NewControllers(service core.IUserService) Controllers {
	c := Controllers{
		service: service,
	}
	return c
}

func (s *Controllers) CreateTenant(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{
			"message": "data validation error",
		})
	} else {
		user.ID = primitive.NewObjectID()
		user.Role = constants.Tenant
		res, err := s.service.CreateUser(c.Request.Context(), &user)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, gin.H{
				"result": res,
			})
		}
	}
}

func (s *Controllers) CreateTenantAdmin(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{
			"message": "data validation error",
		})
	} else {
		user.ID = primitive.NewObjectID()
		user.Role = constants.TenantAdmin
		res, err := s.service.CreateUser(c.Request.Context(), &user)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, gin.H{
				"result": res,
			})
		}
	}
}

func (s *Controllers) GetUser(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		result, err := s.service.GetUserById(c.Request.Context(), id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, result)
		}
	} else {
		email := c.Query("email")
		result, err := s.service.GetUserByEmail(c.Request.Context(), email)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, result)
		}
	}

}

func (s *Controllers) UpdateUser(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{
			"message": "data validation error",
		})
	} else {
		res, err := s.service.UpdateUser(c.Request.Context(), &user)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, gin.H{
				"result": res,
			})
		}
	}
}

func (s *Controllers) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		_, err := s.service.DeleteUser(c.Request.Context(), id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		} else {
			c.Status(204)
		}
	} else {
		c.JSON(400, gin.H{
			"message": "user id not provided",
		})
	}

}
