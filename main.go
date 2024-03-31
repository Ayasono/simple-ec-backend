package main

import (
  database "github.com/Ayasono/simple-kins-backend/database"
  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/Ayasono/simple-kins-backend/routers"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  allowedOrigins := []string{"http://localhost:3000", "http://localhost:8080", "http://127.0.0.1:3000", "http://127.0.0.1:8080"}

  config := cors.Config{
    AllowOrigins:     allowedOrigins,
    AllowHeaders:     []string{"Cache-Control", "Content-Type"},
    AllowCredentials: true,
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
  }

  r.Use(cors.New(config))

  db := database.ConnectDB()

  queries := models.New(db)

  routers.UserRoutes(r, queries)
  routers.ProductRoutes(r, queries)

  r.Run(":8080")
}
