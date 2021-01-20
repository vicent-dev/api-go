package infrastructure

import (
	"api/src/auth"
	"api/src/server/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthFormRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthFormRequest(body io.ReadCloser) (*AuthFormRequest, error) {
	var af AuthFormRequest

	err := json.NewDecoder(body).Decode(&af)
	if err != nil {
		return nil, err
	}
	return &af, nil
}

func SetEndpoints(r *mux.Router) {
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/register", registerHandler).Methods("POST")

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	authForm, err := NewAuthFormRequest(r.Body)

	if err != nil {
		response.WriteErrorResponse(w, err)
		return
	}
	fmt.Printf("%v", authForm)
	auth.NewJWT(nil)
	//get user from mongo

	//generate jwt

	json.NewEncoder(w).Encode(map[string]string{"pata": "patatoe"})
	return
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	authForm, err := NewAuthFormRequest(r.Body)

	if err != nil {
		response.WriteErrorResponse(w, err)
		return
	}
	fmt.Printf("%v", authForm)
	return
}
