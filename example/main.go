package main

import (
	"github.com/mayusys/gorest"
	"net/url"
)

type HelloResource struct{}

func (HelloResource) Get(values url.Values) (int, interface{}) {
	data := map[string]string{"hello": "world"}
	return 200, data
}

func main() {
	helloResource := new(HelloResource)

	var api = new(gorest.API)
	api.AddResource(helloResource, "/hello")
	api.Start(3000)
}
