package main

import (
	"context"
	"server/internal/templates/pages"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Serve static files from the ./static directory
	app.Static("/static", "./server/static")

	// Define a route and handler using the Templ-generated template
	app.Get("/", func(c *fiber.Ctx) error {
		// Create the component using the generated template function
		component := pages.Index("Bradley")

		// Render the component to a string
		var rendered strings.Builder
		ctx := context.Background()
		err := component.Render(ctx, &rendered)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error rendering template")
		}

		// Set the Content-Type header to text/html
		c.Set("Content-Type", "text/html")

		// Send the rendered template as the response
		return c.SendString(rendered.String())
	})

	// Start the server on port 8080
	app.Listen(":8080")
}
