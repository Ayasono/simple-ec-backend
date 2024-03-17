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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowOrigins = append(config.AllowOrigins, "http://localhost:3000")
	config.AllowOrigins = append(config.AllowOrigins, "http://localhost:8080")
	config.AllowOrigins = append(config.AllowOrigins, "http://127.0.0.1:3000")
	config.AllowOrigins = append(config.AllowOrigins, "http://127.0.0.1:8080")
	config.AllowCredentials = true
	r.Use(cors.New(config))

	db := database.ConnectDB()

	queries := models.New(db)

	routers.UserRoutes(r, queries)
	routers.ProductRoutes(r, queries)

	r.Run(":8080")
}
