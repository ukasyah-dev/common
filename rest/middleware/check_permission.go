package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	pb "github.com/ukasyah-dev/pb/authority"
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
