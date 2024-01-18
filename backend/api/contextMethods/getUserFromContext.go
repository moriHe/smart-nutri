package contextmethods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func GetUserFromContext(c *gin.Context) *types.User {
	userInterface, exists := c.Get("user")
	if !exists {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusUnauthorized, Msg: "Not authorized 2"})
		c.Abort()
		return nil
	}

	user, ok := userInterface.(*types.User)

	if !ok || !exists {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusInternalServerError, Msg: "Internal Server Error"})
		c.Abort()
		return nil
	}
	return user
}
