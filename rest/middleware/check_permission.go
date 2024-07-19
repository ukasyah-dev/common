package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ukasyah-dev/common/errors"
)

func CheckPermission(actionID string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return errors.Internal("Not implemented")
	}
}
