package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Controller for routings. using gorrila mux
type Controller struct {
	Router *mux.Router
}

//HTTPFunc type for standart http func
type HTTPFunc func(w http.ResponseWriter, r *http.Request)

//AddRoute add route to router
func (c *Controller) AddRoute(path, method string, fn HTTPFunc) {
	if c.Router == nil {
		c.Router = mux.NewRouter()
	}
	c.Router.HandleFunc(path, fn).Methods(method)
}
