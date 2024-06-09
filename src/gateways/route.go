package gateways

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GatewayHTTP(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("")
	api.Get("/", gateway.IndexPage)
	api.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		gateway.WebSocketServer.HandleWebSocket(ctx)
	}))
	go gateway.WebSocketServer.HandleMessages()
}
