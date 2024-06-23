package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/zeindevs/goshadcn/adapter"
	"github.com/zeindevs/goshadcn/views"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	app := fiber.New()
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(staticFS),
		PathPrefix: "static",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return adapter.FiberRender(ctx, 200, views.Index("go:shadcn/ui"))
	})

	app.Listen(":5000")
}
