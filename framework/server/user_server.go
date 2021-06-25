package server

import (
	"encoding/json"
	"net/http"
	"victorLessa/server/application/usecases"
	"victorLessa/server/domain"
)



type UserServer struct {
	UserUseCase usecases.UserUseCase
}


func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	var c *domain.User
	decoder := json.NewDecoder(r.Body)

	//Aqui ele insere os valores do body no struct da domain user
	if err := decoder.Decode(&c); err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user, err := domain.NewUser(c)

		if err != nil {
			respondWithJSON(w, 400, err)
	}
	res, err := UserServer.UserUseCase.Create(user)
	
	if err != nil {
		respondWithJSON(w, 400, res)
	}
	respondWithJSON(w, 200, res)
}