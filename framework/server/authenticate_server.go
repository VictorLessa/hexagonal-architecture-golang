package server

import (
	"net/http"
	"victorLessa/server/application/usecases"
)


type AuthenticateServer struct {
	AuthenticateUseCases usecases.AuthenticateUseCases
}

func (server *AuthenticateServer) SignIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	res, err := server.AuthenticateUseCases.Authenticate(username, password)

	if err != nil {
		respondWithJSON(w, 401, err)
	}
	respondWithJSON(w, 200, res)
}