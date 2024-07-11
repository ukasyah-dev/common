package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const swaggerHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="SwaggerUI" />
    <title>%s</title>
    <link
      rel="stylesheet"
      href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css"
    />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script
      src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js"
      crossorigin
    ></script>
    <script>
      window.onload = () => {
        window.ui = SwaggerUIBundle({
          url: "./openapi.json",
          dom_id: "#swagger-ui",
        });
      };
    </script>
  </body>
</html>`

func (s *Server) ServeOpenAPISpec() {
	specBytes, _ := s.Config.OpenAPI.Reflector.Spec.MarshalJSON()
	spec := string(specBytes)
	spec = strings.ReplaceAll(spec, "schemas/Model", "schemas/")
	spec = strings.ReplaceAll(spec, "\"Model", "\"")

	s.FiberApp.Get("/docs", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(fmt.Sprintf(swaggerHTML, s.Config.OpenAPI.Spec.Info.Title))
	})

	s.FiberApp.Get("/openapi.json", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.SendString(spec)
	})

}
