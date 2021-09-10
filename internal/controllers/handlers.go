package controllers

import (
	"avito-intern/internal/models"
	"avito-intern/internal/service"
	"avito-intern/internal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	// "io/ioutil"
	"net/http"
)

type MoneyService struct {
	MoneyService service.UserMoneyService
}

func NewMoneyService() *MoneyService {
	return &MoneyService{
		MoneyService: service.NewUserServiceCases(),
	}
}

func (s MoneyService) GetUserMoneyAmount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userUUID := ps.ByName("userId")
	user, err := s.MoneyService.GetUserMoneyAmount(userUUID)

	if err != nil && err == utils.UserNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "get user error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "get user error: %s\n", err)
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "get user json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *MoneyService) CreateTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var tr models.TransactionRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create transaction error: %s\n", err)
		return
	}

	err = json.Unmarshal(body, &tr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "create transaction json decoding error: %s\n", err)
		return
	}

	if tr.Amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "amount must be greater than zero, in request: %v\n", tr.Amount)
		return
	}

	if tr.Amount > 92233720368547758 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong amount: %s\n", utils.NumberIsTooBig)
		return
	}

	tr.UserUUID = ps.ByName("userId")

	createdTransaction, err := s.MoneyService.CreateTransaction(&tr)
	if err != nil && (err == utils.NumberIsTooBig || err == utils.NotEnoughFunds) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "create transaction error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create transaction error: %s\n", err)
		return
	}

	result, err := json.Marshal(createdTransaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "answer for create transaction json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s MoneyService) FundsTransfer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var tr models.TransferRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "transfer error: %s\n", err)
		return
	}

	err = json.Unmarshal(body, &tr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "transfer json decoding error: %s\n", err)
		return
	}

	if tr.Amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "amount must be greater than zero, in request: %v\n", tr.Amount)
		return
	}

	if tr.Amount > 92233720368547758 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong amount: %s\n", utils.NumberIsTooBig)
		return
	}

	createdTransfer, err := s.MoneyService.FundsTransfer(&tr)
	if err != nil && (err == utils.NumberIsTooBig || err == utils.NotEnoughFunds) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "transfer error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "transfer error: %s\n", err)
		return
	}

	result, err := json.Marshal(createdTransfer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "answer for funds transfer json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
