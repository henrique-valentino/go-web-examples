package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/middleware-advanced/http/middleware"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	port := 8081
	fmt.Printf("Port:%d\n", port)
	http.HandleFunc("/", middleware.Chain(Hello, middleware.Method("GET"), middleware.Logging()))
	// http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
