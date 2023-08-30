package types

type RequestError struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func (r *RequestError) Error() string {
	return r.Msg
}
