package restsample

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type APIResource interface {
	Get(values url.Values) (int, interface{})
}

type API struct{}

func (api *API) requestHandler(resource APIResource) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var data interface{}
		var code int

		req.ParseForm()
		methods := req.Method
		value := req.Form

		switch methods {
		case "GET":
			code, data = resource.Get(value)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		content, err := json.Marshal(data)

		if err != nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}

		rw.WriteHeader(code)
		rw.Write(content)
	}
}

func (api *API) AddResource(resource APIResource, path string) {
	http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portString, nil)
}
