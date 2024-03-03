package main

import (
  database "github.com/Ayasono/simple-kins-backend/database"
  models "github.com/Ayasono/simple-kins-backend/models"
  "github.com/Ayasono/simple-kins-backend/routers"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  db := database.ConnectDB()

  queries := models.New(db)

  routers.UserRoutes(r, queries)
  routers.ProductRoutes(r, queries)

  r.Run(":8080")
}
