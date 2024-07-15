package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/errors"
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
