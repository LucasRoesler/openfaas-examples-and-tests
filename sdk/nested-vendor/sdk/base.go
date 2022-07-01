package sdk

import (
	"context"
	"net/http"
)

type Request interface {
	GetBody() []byte
	Headers() http.Header

	GetQueryString() string
	GetMethod() string
	GetHost() string
	Context() context.Context
}

type Response interface {
	GetHeader() http.Header

	GetBody() []byte

	GetStatusCode() int
}

type BasicRequest struct {
	Message string
}

func (r BasicRequest) GetBody() []byte {
	return []byte(r.Message)
}
func (r BasicRequest) Headers() http.Header {
	return nil
}

func (r BasicRequest) GetQueryString() string {
	return ""
}
func (r BasicRequest) GetMethod() string {
	return "GET"
}
func (r BasicRequest) GetHost() string {
	return "localhost"
}

func (r BasicRequest) Context() context.Context {
	return context.TODO()
}
