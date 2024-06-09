package main

import (
	"go-fiber-websocket/configuration"
	gw "go-fiber-websocket/src/gateways"
	"go-fiber-websocket/src/middlewares"
	sv "go-fiber-websocket/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################
	viewsEngine := html.New("./frontend/views", ".html")
	app := fiber.New(configuration.NewFiberConfiguration(viewsEngine))
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	sv0 := sv.NewWebSocket()

	gw.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
