package server

import (
	"crypto"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/emitra-labs/common/errors"
	pb "github.com/emitra-labs/pb/authority"
	"github.com/gofiber/fiber/v2"
	"github.com/swaggest/openapi-go/openapi31"
)

type Server struct {
	FiberApp *fiber.App
	Config   Config
}

type Config struct {
	OpenAPI         OpenAPI
	JWTPublicKey    crypto.PublicKey
	AuthorityClient pb.AuthorityClient
}

type OpenAPI struct {
	Spec      *openapi31.Spec
	Reflector *openapi31.Reflector
}

func New(configs ...Config) *Server {
	config := Config{}
	if len(configs) > 0 {
		config = configs[0]
	}

	if config.OpenAPI.Spec != nil {
		config.OpenAPI.Spec.SetHTTPBearerTokenSecurity("Bearer auth", "JWT", "")

		config.OpenAPI.Reflector = &openapi31.Reflector{
			Spec: config.OpenAPI.Spec,
		}
	}

	return &Server{
		FiberApp: fiber.New(fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError

				if e, ok := err.(*errors.Error); ok {
					code = e.GetHTTPStatus()
				} else if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}

				err = c.Status(code).JSON(fiber.Map{"error": err.Error()})
				if err != nil {
					c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal"})
				}

				return nil
			},
		}),
		Config: config,
	}
}

func (s *Server) Start(port int) error {
	if s.Config.OpenAPI.Reflector != nil {
		s.ServeOpenAPISpec()
	}

	return s.FiberApp.Listen(fmt.Sprintf(":%d", port))
}

func (s *Server) Shutdown() error {
	return s.FiberApp.Shutdown()
}
