package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// returns a 200 with the data appended to the response.
func OK(data interface{}) *Response {
	return NewResponse(
		http.StatusOK,
		"success",
		data,
	)
}

//return a 400 with the error appended as data.
func BadRequest(err error) *Response {
	return NewResponse(
		http.StatusBadRequest,
		"error",
		err.Error(),
	)
}

//returns a 404 with an optional extra message as data.
func NotFound(extraMessages ...string) *Response {
	return NewResponse(
		http.StatusNotFound,
		"not found",
		extraMessages,
	)
}

// return a 500 with the error appended as data.
func ServerError(err error) *Response {
	return ServerErrorWithMessage(err.Error())
}

// return a 500 with a custom message as data.
func ServerErrorWithMessage(message string) *Response {
	return NewResponse(
		http.StatusInternalServerError,
		"error",
		message,
	)
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (r *Response) ToString() string {
	result, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(result)
}
