package services

import (
  "context"
  "net/http"

  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/gin-gonic/gin"
)

// UserService 提供了用户相关的服务。
type UserService struct {
  Queries *models.Queries
}

// ListUsers 返回所有用户。
func (s *UserService) ListUsers(c *gin.Context) {
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
func (s *UserService) CreateUser(c *gin.Context) {
  type CreateUserRequest struct {
    Username     string `json:"username"`
    Email        string `json:"email"`
    PasswordHash string `json:"password_hash"`
  }

  var req CreateUserRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
    return
  }

  user, err := s.Queries.CreateUser(c, models.CreateUserParams{
    Username:     req.Username,
    Email:        req.Email,
    PasswordHash: req.PasswordHash,
  })

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}
