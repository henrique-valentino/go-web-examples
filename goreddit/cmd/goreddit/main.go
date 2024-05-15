package main

import (
	"log"
	"net/http"

	"example.com/goreddit/postgres"
	"example.com/goreddit/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h:= web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
