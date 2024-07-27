package handler

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/emitra-labs/common/rest/middleware"
	"github.com/emitra-labs/common/rest/server"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
)

type Endpoint[I, O any] struct {
	Server      *server.Server
	Method      string
	Path        string
	Func        func(context.Context, *I) (*O, error)
	Summary     string
	Description string
	Tags        []string
}

type Config struct {
	Summary      string
	Description  string
	Tags         []string
	Authenticate bool
	SuperAdmin   bool
	Permission   string
}

func Add[I, O any](srv *server.Server, method string, path string, f func(context.Context, *I) (*O, error), configs ...Config) {
	in := new(I)
	out := new(O)
	funcName := getFuncName(f)

	config := Config{}
	if len(configs) > 0 {
		config = configs[0]
	}

	// Change path params format from :pathParam to {pathParam}
	openapiPath := lo.Reduce(
		strings.Split(path, "/"),
		func(acc string, segment string, i int) string {
			if segment == "" {
				return acc
			}

			if strings.HasPrefix(segment, ":") {
				segment = strings.Replace(segment, ":", "", 1)
				segment = fmt.Sprintf("{%s}", segment)
			}

			return acc + "/" + segment
		},
		"",
	)

	// Add OpenAPI operation
	openapiReflector := srv.Config.OpenAPI.Reflector
	if openapiReflector != nil {
		op, _ := openapiReflector.NewOperationContext(method, openapiPath)
		op.SetID(strcase.ToLowerCamel(funcName))
		op.SetSummary(config.Summary)
		op.SetDescription(config.Description)
		op.SetTags(config.Tags...)
		op.AddReqStructure(in)
		op.AddRespStructure(out)

		if config.Authenticate || config.SuperAdmin || config.Permission != "" {
			op.AddSecurity("Bearer auth")
		}

		if err := openapiReflector.AddOperation(op); err != nil {
			panic(err)
		}
	}

	parseBody := (hasTag(reflect.TypeOf(*in), "json") || hasTag(reflect.TypeOf(*in), "form")) &&
		slices.Contains([]string{"POST", "PUT", "PATCH"}, method)
	parseParams := hasTag(reflect.TypeOf(*in), "params")
	parseQuery := hasTag(reflect.TypeOf(*in), "query")

	fiberHandlers := []fiber.Handler{}

	if config.Authenticate || config.SuperAdmin || config.Permission != "" {
		fiberHandlers = append(fiberHandlers, middleware.Authenticate(srv.Config.JWTPublicKey))
	}

	if config.SuperAdmin {
		fiberHandlers = append(fiberHandlers, middleware.SuperAdmin())
	}

	if config.Permission != "" {
		fiberHandlers = append(fiberHandlers, middleware.CheckPermission(srv.Config.AuthorityClient, config.Permission))
	}

	fiberHandlers = append(fiberHandlers, func(c *fiber.Ctx) error {
		in := new(I)

		if parseBody {
			if err := c.BodyParser(in); err != nil {
				return err
			}
		}

		if parseParams {
			if err := c.ParamsParser(in); err != nil {
				return err
			}
		}

		if parseQuery {
			if err := c.QueryParser(in); err != nil {
				return err
			}
		}

		res, err := f(c.Context(), in)
		if err != nil {
			return err
		}

		return c.JSON(res)
	})

	srv.FiberApp.Add(method, path, fiberHandlers...)
}
