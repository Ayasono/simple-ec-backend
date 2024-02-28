package main

import (
  "net/http"

  "github.com/Ayasono/simple-kins-backend/database"
  database2 "github.com/Ayasono/simple-kins-backend/database/sqlc"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  db := database.ConnectDB()
  queries := database2.New(db)

  r.POST("/users", func(c *gin.Context) {
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

    // 现在直接使用req里的字段，假设数据库操作可以接受string
    user, err := queries.CreateUser(c, database2.CreateUserParams{
      Username:     req.Username,
      Email:        req.Email,
      PasswordHash: req.PasswordHash,
    })

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
  })

  r.GET("/users", func(c *gin.Context) {
    users, err := queries.ListUsers(c)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"users": users})
  })

  r.Run(":8080")
}
