package gateways

import "github.com/gofiber/fiber/v2"

func (h *HTTPGateway) IndexPage(ctx *fiber.Ctx) error {
	data := fiber.Map{"Title": "Hello, World!"}
	return ctx.Render("index", data)
}
