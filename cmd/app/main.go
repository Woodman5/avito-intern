package main

import (
	"avito-intern/internal/config"
	"avito-intern/internal/middleware"
	"avito-intern/internal/routes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	middleware.InitLogger()

	port := fmt.Sprintf(":%v", viper.Get("port"))

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{
		"host": "localhost",
		"port": port,
	}).Info("Starting server")

	router := routes.Routes()

	siteHandler := middleware.AccessLogMiddleware(router)
	siteHandler = middleware.PanicMiddleware(siteHandler)

	log.Fatal(http.ListenAndServe(port, siteHandler))
}
