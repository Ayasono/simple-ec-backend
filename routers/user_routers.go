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
    userGroup.GET("/", func(context *gin.Context) {
      services.ListUsers(context, queries)
    })

    userGroup.POST("/", func(context *gin.Context) {
      services.CreateUser(context, queries)
    })

    userGroup.GET("/:email", func(context *gin.Context) {
      services.GetUserByEmail(context, queries)
    })
  }
}
