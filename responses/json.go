package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BodyError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type Data struct {
	UserId  uint32 `json:"userId"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Token   string `json:"token"`
}

func Json(w http.ResponseWriter, statusCode int, data Data) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
}

func Error(w http.ResponseWriter, statusCode int, err string) bool {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp := BodyError{
		Message: err,
		Code:    statusCode,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
	return true
}
