package main

import (
	"log"
	"net/http"

	"git.v6.io/go-web/middleware"
	"git.v6.io/go-web/routes"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/item/{id}", routes.FindItemById)
	router.HandleFunc("/item/latest", routes.FindItemLatest)
	router.HandleFunc("POST /items", routes.CreateItem)

	stack := middleware.CreateStack(middleware.Logging, middleware.SampleMiddleware)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(v1),
		// Handler: stack(router),
		// Handler: router,
	}

	log.Println("Server started on port :8080")
	server.ListenAndServe()
}
