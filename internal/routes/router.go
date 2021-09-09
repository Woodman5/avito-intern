package routes

import (
	"avito-intern/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	service := controllers.NewUserPetService()

	router := httprouter.New()
	router.POST("/amount/:userId", service.CreateTransaction)
	router.GET("/amount/:userId", service.GetUserMoney)
	router.GET("/history/:userId", service.GetUserHistory)

	return router
}
