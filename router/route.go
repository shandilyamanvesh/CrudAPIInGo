package router

import (
	"net/http"
)

// Structure of a particular route.
type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type routes []route
