package http

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request) (status int, err error)
