package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"lemonilo.app/auth"
	"lemonilo.app/database"
	"lemonilo.app/model"
	response "lemonilo.app/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	resp := response.Data{
		Token: token,
	}

	response.Json(w, http.StatusOK, resp)
}

func SignIn(email, password string) (string, error) {
	var err error

	user := model.User{}

	err = database.Connector.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.UserId)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
