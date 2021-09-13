package routes

import (
	"victorLessa/server/application/repositories"
	"victorLessa/server/application/usecases"
	"victorLessa/server/framework/server"

	mux "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


type UserRoute struct {
	Router *mux.Router
	Db *gorm.DB
}

func (UserRoute *UserRoute) UserRoutes() {

	userRepository := repositories.UserRepositoryDb{Db:  UserRoute.Db}
	userUsecases := usecases.UserUseCase{UserRepository: userRepository}
	userServer := server.UserServer{UserUseCase: userUsecases}

	UserRoute.Router.HandleFunc("/users", userServer.CreateUser).Methods("POST")
	UserRoute.Router.HandleFunc("/users", userServer.IndexUsers).Methods("GET")
	UserRoute.Router.HandleFunc("/users/{id}", userServer.ShowUsers).Methods("GET")
	UserRoute.Router.HandleFunc("/users/{id}", userServer.UpdateUser).Methods("PUT")
	UserRoute.Router.HandleFunc("/users/{id}", userServer.UpdateUser).Methods("PUT")
	UserRoute.Router.HandleFunc("/users/{id}", userServer.DeleteUser).Methods("DELETE")
}
