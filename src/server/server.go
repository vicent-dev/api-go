package server

import (
	"net/http"

	"api/src/server/response"
	user "api/src/user/infrastructure"

	"github.com/gorilla/mux"
)

func Start(port string) {

	r := mux.NewRouter()
	r.Use(response.JsonMiddleware)

	user.SetEndpoints(r)

	http.ListenAndServe(":"+port, r)
}
