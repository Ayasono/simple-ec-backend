package services

import (
	"context"

	models "github.com/Ayasono/simple-kins-backend/models"
	"github.com/gin-gonic/gin"
)

func (s *Services) ListProducts(c *gin.Context) {
	products, err := s.Queries.ListProducts(context.Background(), models.ListProductsParams{
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

func (s *Services) GetProductCategories(c *gin.Context) {
	categories, err := s.Queries.GetProductCategories(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"msg":        "ok",
		"categories": categories,
	})
}
