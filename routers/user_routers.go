package routers

import (
  database "github.com/Ayasono/simple-kins-backend/models"
  "github.com/Ayasono/simple-kins-backend/services"
  "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, queries *database.Queries) {
  userGroup := router.Group("/users")
  {
    userGroup.GET("/", (&services.UserService{Queries: queries}).ListUsers)

    userGroup.POST("/", (&services.UserService{Queries: queries}).CreateUser)
  }
}
