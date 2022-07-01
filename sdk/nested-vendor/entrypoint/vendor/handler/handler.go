package handler

import (
	"net/http"

	"github.com/LucasRoesler/openfaas-examples-and-tests/sdk/nested-vendor/sdk"
)

func Handle(sdk.Request) (sdk.Response, error) {
	return resp{msg: "hi", code: 200}, nil
}

type resp struct {
	msg  string
	code int
}

func (r resp) GetHeader() http.Header {
	return nil
}

func (r resp) GetBody() []byte {
	return []byte(r.msg)
}

func (r resp) GetStatusCode() int {
	return r.code
}
