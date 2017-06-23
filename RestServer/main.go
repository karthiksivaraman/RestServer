package main

import (
	"net/http"

	"github.com/go-zoo/bone"
)

func main() {
	mux := bone.New()

	// mux.Get, Post, etc ... takes http.Handler
	mux.Get("/home/:id", http.HandlerFunc(HomeHandler))

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	// Get the value of the "id" parameters.
	val := bone.GetValue(req, "id")

	rw.Write([]byte(val))
}
