package controllers

import (
	// "avito-intern/internal/models"
	"avito-intern/internal/service"
	"avito-intern/internal/utils"
	"encoding/json"
	"fmt"

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

// func (s *UserPetService) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	var user models.User

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "create user error: %s\n", err)
// 		return
// 	}

// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "create user json decoding error: %s\n", err)
// 		return
// 	}

// 	userUUID, err := uuid.NewUUID()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "create user UUID generation error: %s\n", err)
// 		return
// 	}

// 	user.UUID = userUUID.String()

// 	createdUser, err := s.UserService.CreateUser(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "create user error: %s\n", err)
// 		return
// 	}

// 	result, err := json.Marshal(createdUser)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "create user json encoding error: %s\n", err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)
// }

// func (s *UserPetService) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	var user models.User

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "update user error: %s\n", err)
// 		return
// 	}

// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "update user json decoding error: %s\n", err)
// 		return
// 	}

// 	updatedUser, err := s.UserService.UpdateUser(&user)
// 	if err != nil && err == utils.UserNotFound {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "update user error: %s\n", err)
// 		return
// 	} else if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "update user error: %s\n", err)
// 		return
// 	}

// 	result, err := json.Marshal(updatedUser)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "update user json encoding error: %s\n", err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)
// }

// func (s *UserPetService) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	userUUID := ps.ByName("userId")
// 	_, err := s.UserService.DeleteUser(userUUID)

// 	if err != nil && err == utils.UserNotFound {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintf(w, "delete user error: %s\n", err)
// 		return
// 	} else if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "delete user error: %s\n", err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, `{"deleted_uuid":"%s"}`, userUUID)

// }
