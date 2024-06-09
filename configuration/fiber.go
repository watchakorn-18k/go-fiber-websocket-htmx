package configuration

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewFiberConfiguration(viewsEngine *html.Engine) fiber.Config {
	return fiber.Config{
		AppName:     ")϶ go-fiber-websocket ϵ(",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Views:       viewsEngine,
	}
}
