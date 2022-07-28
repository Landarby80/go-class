package main

import (
	"chat/router"
	"net/http"

	// "os"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	router := router.ConfigureRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	addr := ":8000"

	logrus.WithField("addr", addr).Info("Starting server...")

	if err := http.ListenAndServe(addr, handlers.CORS(headers, methods, origins)(router)); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}

}
