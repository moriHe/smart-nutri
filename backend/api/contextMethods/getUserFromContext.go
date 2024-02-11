package contextmethods

import (
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func GetUserFromContext(c *gin.Context) *types.User {
	userInterface, exists := c.Get("user")
	if !exists {
		responses.ErrorResponse(c, &types.UnauthorizedError)
		return nil
	}

	user, ok := userInterface.(*types.User)

	if !ok || !exists {
		responses.ErrorResponse(c, &types.InternalServerError)
		return nil
	}
	return user
}
