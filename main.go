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

  r.Use(cors.New(cors.Config{
    AllowOrigins:  []string{"*"},
    AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
    AllowHeaders:  []string{"Origin", "Content-Type"},
    ExposeHeaders: []string{"Content-Length"},
  }))

  db := database.ConnectDB()

  queries := models.New(db)

  routers.UserRoutes(r, queries)
  routers.ProductRoutes(r, queries)

  r.Run(":8080")
}
