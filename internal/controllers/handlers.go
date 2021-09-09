package controllers

import (
	"avito-intern/internal/models"
	"avito-intern/internal/service"
	"avito-intern/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type UserPetService struct {
	UserService service.UserService
}

func NewUserPetService() *UserPetService {
	return &UserPetService{
		UserService: service.NewUserService(),
	}
}

func (s UserPetService) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userUUID := ps.ByName("userId")
	user, err := s.UserService.GetUserByUUID(userUUID)

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

func (s UserPetService) GetPet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userUUID := ps.ByName("userId")
	petUUID := ps.ByName("petId")
	pet, err := s.UserService.GetPetByUUID(userUUID, petUUID)

	if err != nil && (err == utils.PetNotFound || err == utils.UserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "get pet error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "get pet error: %s\n", err)
		return
	}

	result, err := json.Marshal(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "get pet json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *UserPetService) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create user error: %s\n", err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "create user json decoding error: %s\n", err)
		return
	}

	userUUID, err := uuid.NewUUID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create user UUID generation error: %s\n", err)
		return
	}

	user.UUID = userUUID.String()

	createdUser, err := s.UserService.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create user error: %s\n", err)
		return
	}

	result, err := json.Marshal(createdUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create user json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *UserPetService) CreatePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create pet error: %s\n", err)
		return
	}

	var pet models.Pet

	err = json.Unmarshal(body, &pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "create pet json decoding error: %s\n", err)
		return
	}

	petUUID, err := uuid.NewUUID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create pet UUID generation error: %s\n", err)
		return
	}

	pet.UUID = petUUID.String()

	userUUID := ps.ByName("userId")

	createdPet, err := s.UserService.CreatePet(userUUID, &pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create pet error: %s\n", err)
		return
	}

	result, err := json.Marshal(createdPet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "create pet json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *UserPetService) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update user error: %s\n", err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "update user json decoding error: %s\n", err)
		return
	}

	updatedUser, err := s.UserService.UpdateUser(&user)
	if err != nil && err == utils.UserNotFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "update user error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update user error: %s\n", err)
		return
	}

	result, err := json.Marshal(updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update user json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *UserPetService) UpdatePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update pet error: %s\n", err)
		return
	}

	var pet models.Pet

	err = json.Unmarshal(body, &pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "update pet json decoding error: %s\n", err)
		return
	}

	userUUID := ps.ByName("userId")

	updatedPet, err := s.UserService.UpdatePet(userUUID, &pet)
	if err != nil && (err == utils.PetNotFound || err == utils.UserNotFound) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "update pet error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update pet error: %s\n", err)
		return
	}

	result, err := json.Marshal(updatedPet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update pet json encoding error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (s *UserPetService) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userUUID := ps.ByName("userId")
	_, err := s.UserService.DeleteUser(userUUID)

	if err != nil && err == utils.UserNotFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "delete user error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "delete user error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"deleted_uuid":"%s"}`, userUUID)

}

func (s *UserPetService) DeletePet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userUUID := ps.ByName("userId")
	petUUID := ps.ByName("petId")
	_, err := s.UserService.DeletePet(userUUID, petUUID)

	if err != nil && (err == utils.PetNotFound || err == utils.UserNotFound) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "delete pet error: %s\n", err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "delete pet error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"deleted_uuid":"%s"}`, petUUID)

}
