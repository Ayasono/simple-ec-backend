package services

import (
	"context"
	"fmt"

	models "github.com/Ayasono/simple-kins-backend/models"
	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context, queries *models.Queries) {
	// 假设你想缓存这个响应1小时
	maxAge := 3600 // 单位为秒

	products, err := queries.ListProducts(context.Background(), models.ListProductsParams{
		Limit:  50,
		Offset: 0,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 设置 Cache-Control 头部，允许客户端和代理服务器缓存这个响应1小时
	c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))

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
