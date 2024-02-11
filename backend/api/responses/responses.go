package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

type ErrorDetails struct {
	Message string `json:"message"`
}
type Response struct {
	Error        bool          `json:"error"`
	Status       int           `json:"status"`
	Data         interface{}   `json:"data,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
}

func ErrorResponse(c *gin.Context, err *types.RequestError) {
	c.JSON(err.Status, Response{
		Error:        true,
		Status:       err.Status,
		ErrorDetails: &ErrorDetails{Message: err.Msg},
	})
}

func HandleResponse(c *gin.Context, data any, err *types.RequestError) {
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Error:  false,
		Status: http.StatusOK,
		Data:   data,
	})
}
