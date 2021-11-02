package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// ROUTER SETUP
	r := chi.NewRouter()
	
	// Logs start/end, response status, and more of each request
	r.Use(middleware.Logger)
	// Timeout (60 second timeout)
	r.Use(middleware.Timeout(60*time.Second))
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}