package routes

import (
	"avito-intern/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	service := controllers.NewMoneyService()

	router := httprouter.New()
	router.POST("/amount/:userId", service.CreateTransaction)
	router.POST("/transfer", service.FundsTransfer)
	router.GET("/amount/:userId", service.GetUserMoneyAmount)
	//router.GET("/history/:userId", service.FundsTransfer)

	return router
}
