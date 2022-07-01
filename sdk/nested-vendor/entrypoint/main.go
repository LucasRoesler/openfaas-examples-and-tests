package main

import (
	"fmt"
	"handler"

	"github.com/LucasRoesler/openfaas-examples-and-tests/sdk/nested-vendor/sdk"
)

func main() {
	fmt.Println(GetMessage(handler.Handle))
}

func GetMessage(f func(sdk.Request) (sdk.Response, error)) string {
	resp, err := f(sdk.BasicRequest{Message: "hola"})
	if err != nil {
		panic(err)
	}

	return string(resp.GetBody())
}
