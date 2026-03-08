package middleware

import(
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlwares ...Middleware) http.Handler {
	for i := len(middlwares) -1; i >= 0; i-- {
		h = middlwares[i](h)
	}
	return h
}