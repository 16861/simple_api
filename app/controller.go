package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	Router *mux.Router
}

type HTTPFunc func(w http.ResponseWriter, r *http.Request)

func (c *Controller) AddRoute(path, method string, fn HTTPFunc) {
	if c.Router == nil {
		c.Router = mux.NewRouter()
	}
	c.Router.HandleFunc(path, fn).Methods(method)
}
