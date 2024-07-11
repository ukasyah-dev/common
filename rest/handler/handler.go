package handler

import (
	"context"
	"reflect"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
	"github.com/ukasyah-dev/common/rest/server"
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
	Summary     string
	Description string
	Tags        []string
}

func Add[I, O any](srv *server.Server, method string, path string, f func(context.Context, *I) (*O, error), configs ...Config) {
	in := new(I)
	out := new(O)
	funcName := getFuncName(f)

	config := Config{}
	if len(configs) > 0 {
		config = configs[0]
	}

	// Add OpenAPI operation
	openapiReflector := srv.Config.OpenAPI.Reflector
	if openapiReflector != nil {
		op, _ := openapiReflector.NewOperationContext(method, path)
		op.SetID(strcase.ToLowerCamel(funcName))
		op.SetSummary(config.Summary)
		op.SetDescription(config.Description)
		op.SetTags(config.Tags...)
		op.AddReqStructure(in)
		op.AddRespStructure(out)
		openapiReflector.AddOperation(op)
	}

	parseBody := hasTag(reflect.TypeOf(*in), "json") && slices.Contains([]string{"POST", "PUT", "PATCH"}, method)
	parseParams := hasTag(reflect.TypeOf(*in), "params")
	parseQuery := hasTag(reflect.TypeOf(*in), "query")

	fiberHandlers := []fiber.Handler{}

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
