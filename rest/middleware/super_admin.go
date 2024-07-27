package middleware

import (
	"github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/gofiber/fiber/v2"
)

func SuperAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		superAdmin, ok := c.Locals(constant.SuperAdmin).(bool)
		if !ok || !superAdmin {
			return errors.PermissionDenied()
		}

		return c.Next()
	}
}
