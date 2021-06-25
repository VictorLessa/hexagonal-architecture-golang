package main

import (
	"encoding/json"
	"net/http"
	"os"
	"victorLessa/server/application/repositories"
	"victorLessa/server/application/usecases"
	"victorLessa/server/framework/database"
	"victorLessa/server/framework/server"

	mux "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


type App struct {
	Router *mux.Router
	Db *gorm.DB
}

func (a *App) Initialize() {
	
	a.Db = database.ConnectDb(os.Getenv("env"))
	
	a.Router = mux.NewRouter()
	
	a.initializeRoutes()
}


func (a * App) initializeRoutes() {
	a.Router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		response, _ := json.Marshal("Hello World")
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)

	}).Methods("GET")

	userRepository := repositories.UserRepositoryDb{Db:  a.Db}
	userUsecases := usecases.UserUseCase{UserRepository: userRepository}	
	userServer := server.UserServer{UserUseCase: userUsecases}

	a.Router.HandleFunc("/users", userServer.CreateUser).Methods("POST")
}

func (a *App) Run() {
	http.ListenAndServe(":8010", a.Router)

}


func main() {
	a := App{}

	a.Initialize()
	a.Run()
	
}