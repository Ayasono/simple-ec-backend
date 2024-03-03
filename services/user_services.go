package services

import (
  "context"
  "net/http"

  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/gin-gonic/gin"
)

// ListUsers 返回所有用户。
func ListUsers(c *gin.Context, queries *models.Queries) {
  users, err := queries.ListUsers(context.Background())
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  }
  // 成功情况下，返回用户列表和200 OK状态码
  c.JSON(http.StatusOK, users)
}

// CreateUser 是创建用户的请求体。
func CreateUser(c *gin.Context, queries *models.Queries) {
  type userRequest struct {
    Username     string `json:"username" binding:"required"`
    Email        string `json:"email" binding:"required"`
    PasswordHash string `json:"password_hash" binding:"required"`
  }

  var req userRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  // check if email is already in use before creating a new user
  email, err := queries.GetUserByEmail(context.Background(), req.Email)
  if err == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": email.Email + " is already in use"})
    return
  }
  // 创建用户
  userInfo, err := queries.CreateUser(context.Background(), models.CreateUserParams{
    Username:     req.Username,
    Email:        req.Email,
    PasswordHash: req.PasswordHash,
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
    return
  }
  // 返回用户ID
  c.JSON(http.StatusOK, gin.H{
    "msg":  "User created",
    "user": userInfo,
  })
}

// GetUserByEmail 是获取用户的请求体。
func GetUserByEmail(c *gin.Context, queries *models.Queries) {
  email := c.Param("email")
  user, err := queries.GetUserByEmail(context.Background(), email)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg":  "ok",
    "user": user,
  })
}
