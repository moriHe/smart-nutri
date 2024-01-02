package contextmethods

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func GetUserFromContext(c *gin.Context) (*types.User, error) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, errors.New("Not authorized")
	}

	user, ok := userInterface.(*types.User)

	if !ok {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusInternalServerError, Msg: "Internal Server Error"})
		return nil, errors.New("Not authorized")
	}
	return user, nil
}
