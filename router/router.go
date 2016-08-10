package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	muxrouter := mux.NewRouter().StrictSlash(true)
	for _, route := range muxroutes {

		var handler http.Handler
		handler = route.handlerFunc

		muxrouter.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(handler)
	}
	return muxrouter
}
