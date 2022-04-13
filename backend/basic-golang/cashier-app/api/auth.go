package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	res, err := api.usersRepo.Login(username, password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	fmt.Println(res)

<<<<<<< HEAD
	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: username}) // TODO: replace this
=======
	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: ""}) // TODO: replace this
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	err := api.usersRepo.Logout(username)
<<<<<<< HEAD
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
=======
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c

	encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this
}
