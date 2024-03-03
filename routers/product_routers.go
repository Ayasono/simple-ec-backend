package routers

import (
  database "github.com/Ayasono/simple-kins-backend/models"
  "github.com/Ayasono/simple-kins-backend/services"
  "github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, queries *database.Queries) {
  // product group
  productGroup := router.Group("/products")
  {
    productGroup.GET("/", func(context *gin.Context) {
      services.ListProducts(context, queries)
    })

    productGroup.GET("/categories", func(context *gin.Context) {
      services.GetProductCategories(context, queries)
    })
  }
}
