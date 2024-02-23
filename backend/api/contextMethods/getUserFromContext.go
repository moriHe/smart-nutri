package contextmethods

import (
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func GetUserFromContext(c *gin.Context) *types.User {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil
	}

	user, ok := userInterface.(*types.User)

	if !ok || !exists {
		return nil
	}
	return user
}
