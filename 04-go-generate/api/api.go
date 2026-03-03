package api

import (
	"log"
	"net/http"

	_ "go-generate/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Simple Go API
// @version 1.0
// @description This is a simple API built with Go and documented with Swagger.
// @host localhost:8090
// @BasePath /
func StartServer() error {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(HelloWorld))
	mux.HandleFunc("/openapi/", httpSwagger.Handler())

	log.Println("Server is running on http://localhost:8090")
	return http.ListenAndServe("0.0.0.0:8090", mux)
}

// @Summary Hello World Endpoint
// @Description Responds with "Hello World"
// @Tags example
// @Produce text/html
// @Success 200 {string} string "Hello World"
// @Router /hello [get]
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
