package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

/*
Function for initializing new router instance.
Also takes routes array and maps this with respective handlers.
*/

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
