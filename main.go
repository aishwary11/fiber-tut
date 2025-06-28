package main

import (
	"log"
	"os"

	"github.com/aishwary11/fiber-tut/middleware"
	"github.com/aishwary11/fiber-tut/routes"
	"github.com/aishwary11/fiber-tut/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.ConnectDB()
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// kafka.InitProducer("localhost:9092", "video-events")
	// defer kafka.CloseProducer()
	// kafka.StartConsumer("localhost:9092", "video-events", "my-group")
	// kafka.ProduceMessage("video:123", "play_started")
	app.Use(middleware.Logger(port))
	app.Use(middleware.RateLimit())
	routes.UserRoutes(app)
	app.Use(middleware.JWTMiddleware())

	app.Listen(":" + port)
}
