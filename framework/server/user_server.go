package server

import (
	"encoding/json"
	"net/http"
	"victorLessa/server/application/usecases"
	"victorLessa/server/domain"

	"github.com/gorilla/mux"
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
		respondWithJSON(w, 400, err)
	}
	respondWithJSON(w, 200, res)
}

func (UserServer *UserServer) IndexUsers(w http.ResponseWriter, r *http.Request) {

	res := UserServer.UserUseCase.Index()

	respondWithJSON(w, 200, res.Value)
}

func (UserServer *UserServer) ShowUsers(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res := UserServer.UserUseCase.Show(id)

	respondWithJSON(w, 200, res.Value)

}

func (UserServer *UserServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var c *domain.User
	id := mux.Vars(r)["id"]
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

	res, err := UserServer.UserUseCase.Update(id, user)

	if err != nil {
		respondWithJSON(w, 400, err)
	}

	respondWithJSON(w, 200, res)
}

func (UserServer *UserServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	res, err := UserServer.UserUseCase.Delete(id)

	if err != nil {
		respondWithJSON(w, 400, err)
	}

	respondWithJSON(w, 200, res)

}