package main

import (
	"avito-intern/internal/config"
	"avito-intern/internal/middleware"
	"avito-intern/internal/routes"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

const Port = ":8080"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	middleware.InitLogger()

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{
		"host": "localhost",
		"port": Port,
	}).Info("Starting server")

	router := routes.Routes()

	siteHandler := middleware.AccessLogMiddleware(router)
	siteHandler = middleware.PanicMiddleware(siteHandler)

	log.Fatal(http.ListenAndServe(Port, siteHandler))
}
