package main

import (
  "database/sql"
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
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // 转换 CreateUserRequest 为 CreateUserParams
    params := database2.CreateUserParams{
      Username:     sql.NullString{String: req.Username, Valid: true},
      Email:        sql.NullString{String: req.Email, Valid: true},
      PasswordHash: sql.NullString{String: req.PasswordHash, Valid: true},
    }

    user, err := queries.CreateUser(c, params)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
  })

  r.Run(":8080")
}
