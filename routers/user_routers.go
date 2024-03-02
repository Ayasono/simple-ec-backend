package routers

import (
	database "github.com/Ayasono/simple-kins-backend/models"
	"github.com/Ayasono/simple-kins-backend/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, queries *database.Queries) {
	// users group
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", (&services.Services{Queries: queries}).ListUsers)

		userGroup.POST("/", (&services.Services{Queries: queries}).CreateUser)

		userGroup.GET("/:email", (&services.Services{Queries: queries}).GetUserByEmail)
	}
	// product group
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", (&services.Services{Queries: queries}).ListProducts)
		productGroup.GET("/categories", (&services.Services{Queries: queries}).GetProductCategories)
	}
}
