package gateways

import (
	service "go-fiber-websocket/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	WebSocketServer service.IWebSocketService
}

func NewHTTPGateway(app *fiber.App, websocket service.IWebSocketService) {
	gateway := &HTTPGateway{
		WebSocketServer: websocket,
	}

	GatewayHTTP(*gateway, app)
}
