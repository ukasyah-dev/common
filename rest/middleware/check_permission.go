package middleware

import (
	"github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	pb "github.com/emitra-labs/pb/authority"
	"github.com/gofiber/fiber/v2"
)

func CheckPermission(authorityClient pb.AuthorityClient, actionID string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		teamID := c.Params("teamId")
		if teamID == "" {
			return errors.InvalidArgument("Missing team ID")
		}

		userID, ok := c.Locals(constant.UserID).(string)
		if !ok || userID == "" {
			return errors.InvalidArgument("Missing user ID")
		}

		res, err := authorityClient.CheckPermission(c.Context(), &pb.CheckPermissionRequest{
			ActionID: actionID,
			TeamID:   teamID,
			UserID:   userID,
		})
		if err != nil {
			log.Errorf("Failed to check permission: %s", err)
			return errors.Internal()
		}

		if res.Allowed {
			return c.Next()
		}

		return errors.PermissionDenied()
	}
}
