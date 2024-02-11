package types

type RequestError struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func (r *RequestError) Error() string {
	return r.Msg
}

var UnauthorizedError RequestError = RequestError{Status: 401, Msg: "User not authorized."}
var BadRequestError RequestError = RequestError{Status: 400, Msg: "Bad request."}
var InternalServerError RequestError = RequestError{Status: 500, Msg: "Internal Server Error."}

func NewRequestError(errorType *RequestError, message string) error {
	requestError := errorType
	requestError.Msg = message
	return requestError
}
