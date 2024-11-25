package main

import (
	"net/http"
	"runtime/debug"
)

func (log *applog) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	log.Logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// func (log *applog) clientError(w http.ResponseWriter, status int) {
//     http.Error(w, http.StatusText(status), status)
// }
