package services

import (
  "context"

  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context, queries *models.Queries) {
  products, err := queries.ListProducts(context.Background(), models.ListProductsParams{
    Limit:  50,
    Offset: 0,
  })
  if err != nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
    "msg":      "ok",
    "products": products,
  })
}

func GetProductCategories(c *gin.Context, queries *models.Queries) {
  categories, err := queries.GetProductCategories(context.Background())
  if err != nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
    "msg":        "ok",
    "categories": categories,
  })
}
