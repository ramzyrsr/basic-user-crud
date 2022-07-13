package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"lemonilo.app/database"
	"lemonilo.app/helper"
	"lemonilo.app/model"
	response "lemonilo.app/responses"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var persons []model.User

	getUsers := database.Connector.Find(&persons)
	if getUsers.Error != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.JSON(w, http.StatusOK, &getUsers)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["userId"]

	person := model.User{}
	getUser := database.Connector.First(&person, "user_id = ?", key)
	if getUser.Error != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	resp := response.Data{
		UserId:  person.UserId,
		Email:   person.Email,
		Address: person.Address,
	}

	response.Json(w, http.StatusOK, resp)
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	person := model.User{}
	err = json.Unmarshal(requestBody, &person)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	// validate required request
	errVal := helper.SendValidation(&person)
	if errVal != "" {
		response.Error(w, http.StatusBadRequest, "Field "+errVal+"cannot be empty")
		return
	}

	person.UserId = uuid.New().ID()

	hashPassword, err := Hash(person.Password)
	if err != nil {
		return
	}
	person.Password = string(hashPassword)

	createUser := database.Connector.Create(&person)
	if createUser.Error != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	resp := response.Data{
		UserId:  person.UserId,
		Email:   person.Email,
		Address: person.Address,
	}

	response.Json(w, http.StatusCreated, resp)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["userId"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	person := model.User{}
	err = json.Unmarshal(requestBody, &person)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	updateUser := database.Connector.Model(&person).Where("user_id = ?", key).Update(&person).Take(&person)
	if updateUser.Error != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Internal Server Error")
		return
	}

	resp := response.Data{
		UserId:  person.UserId,
		Email:   person.Email,
		Address: person.Address,
	}

	response.Json(w, http.StatusOK, resp)
}

func DeletUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["userId"]

	person := model.User{}
	deleteUser := database.Connector.Where("user_id = ?", key).Delete(&person)
	if deleteUser.Error != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Internal Server Error")
		return
	}

	resp := response.Data{
		UserId:  person.UserId,
		Email:   person.Email,
		Address: person.Address,
	}

	response.Json(w, http.StatusNoContent, resp)
}
