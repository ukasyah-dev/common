package middleware

import (
	"crypto"
	"strings"

	"github.com/emitra-labs/common/auth"
	"github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(publicKey crypto.PublicKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.Replace(c.Get("Authorization"), "Bearer ", "", 1)
		if accessToken == "" {
			return errors.Unauthenticated()
		}

		claims, err := auth.VerifyAccessToken(publicKey, accessToken)
		if err != nil {
			return errors.Unauthenticated()
		}

		c.Locals(constant.UserID, claims.UserID)
		c.Locals(constant.SessionID, claims.SessionID)
		c.Locals(constant.SuperAdmin, claims.SuperAdmin)

		return c.Next()
	}
}
