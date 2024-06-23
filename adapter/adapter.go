package adapter

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func FiberRender(ctx *fiber.Ctx, status int, template templ.Component) error {
	ctx.Status(status)
	handler := adaptor.HTTPHandler(templ.Handler(template))
	return handler(ctx)
}
