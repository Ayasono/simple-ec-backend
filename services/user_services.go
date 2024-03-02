package services

import (
  "context"
  "net/http"

  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/gin-gonic/gin"
)

// UserService 提供了用户相关的服务。

// ListUsers 返回所有用户。
func (s *Services) ListUsers(c *gin.Context) {
  users, err := s.Queries.ListUsers(context.Background())
  if err != nil {
    // 处理错误的情况，这里简单返回500错误
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
    return
  }
  // 成功情况下，返回用户列表和200 OK状态码
  c.JSON(http.StatusOK, users)
}

// CreateUser 是创建用户的请求体。
func (s *Services) CreateUser(c *gin.Context) {
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
  email, err := s.Queries.GetUserByEmail(context.Background(), req.Email)
  if err == nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": email.Email + " is already in use"})
    return
  }
  // 创建用户
  userInfo, err := s.Queries.CreateUser(context.Background(), models.CreateUserParams{
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
func (s *Services) GetUserByEmail(c *gin.Context) {
  email := c.Param("email")
  user, err := s.Queries.GetUserByEmail(context.Background(), email)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg":  "ok",
    "user": user,
  })
}
