package handler

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New()

	app.Get("/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"version": "v1",
		})
	})
	app.Get("/v2", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"version": "v2",
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"url":  c.Request().URI().String(),
			"path": c.Path(),
		})
	})

	return adaptor.FiberApp(app)
}
