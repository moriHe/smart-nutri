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

func ErrorResponse(c *gin.Context, err error) {
	if requestErr, ok := err.(*types.RequestError); ok {
		c.JSON(requestErr.Status, Response{
			Error:        true,
			Status:       requestErr.Status,
			ErrorDetails: &ErrorDetails{Message: requestErr.Msg},
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, Response{
			Error:        true,
			Status:       http.StatusInternalServerError,
			ErrorDetails: &ErrorDetails{Message: "Unhandled error type"},
		})
	}
}

func HandleResponse(c *gin.Context, data any, err error) {
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
