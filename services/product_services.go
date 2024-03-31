package services

import (
  "context"
  "fmt"
  "net/http"
  "strconv"

  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/Ayasono/simple-kins-backend/utils"
  "github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context, queries *models.Queries) {
  maxAge := 3600 // 单位为秒
  const defaultPage = 0
  const defaultPageSize = 50

  // 尝试转换 page，如果失败则使用默认值
  page, err := strconv.Atoi(c.Query("page")) // 从 query string 中获取 page 参数
  if err != nil || page < 1 {
    page = defaultPage
  }

  // 尝试转换 pageSize，如果失败则使用默认值
  pageSize, err := strconv.Atoi(c.Query("pageSize"))
  if err != nil || pageSize <= 0 {
    pageSize = defaultPageSize
  }

  offset := page * pageSize

  products, err := queries.ListProducts(context.Background(), models.ListProductsParams{
    Limit:  int32(pageSize),
    Offset: int32(offset),
  })

  if err != nil {
    c.JSON(http.StatusInternalServerError, utils.GeneralResStruct{
      Msg:   "not ok",
      Error: err.Error(),
    })
    return
  }

  // 设置 Cache-Control 头部，允许客户端和代理服务器缓存这个响应1小时
  c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))

  c.JSON(http.StatusOK, utils.GeneralResStruct{
    Msg:  "ok",
    Data: products,
  })
}

func GetProductCategories(c *gin.Context, queries *models.Queries) {
  categories, err := queries.GetProductCategories(context.Background())
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg":        "ok",
    "categories": categories,
  })
}
