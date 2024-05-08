package middleware

import "net/http"

type Middleware func (http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, m := range middlewares{
		m(f)
	}
	return f
}