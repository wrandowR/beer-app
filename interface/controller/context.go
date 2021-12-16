package controller

import (
	"io"
	"mime/multipart"
	"net/http"
)

//Context context for api rest
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	QueryParam(name string) string
	String(code int, s string) error
	Get(key string) interface{}
	FormFile(name string) (*multipart.FileHeader, error)
	Stream(code int, contentType string, r io.Reader) error
	NoContent(code int) error
	Request() *http.Request
}
