package logger

import (
	"log"
	"net/http"
	"time"
)

// Monkey patching http handler with a functionality that logs basic details of every request on terminal.
func Logger(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
