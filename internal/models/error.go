package models

import (
	"net/http"

	"github.com/go-chi/render"
)

const ParseErr string = "Can't parse "
const GetListErr string = "Can't get list"
const GetItemErr string = "Can't get item"
const NotFoundErr string = "Not Found"
const BadReqErr string = "Bad request"

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error, message string, code int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: code,
		StatusText:     message,
		ErrorText:      err.Error(),
	}
}
