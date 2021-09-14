package routes

import (
	"victorLessa/server/application/middleware"
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

	userRoute := UserRoute.Router 
	userRoute.HandleFunc("/users", userServer.CreateUser).Methods("POST")
	userRoute.HandleFunc("/users", userServer.IndexUsers).Methods("GET")
	userRoute.HandleFunc("/users/{id}", userServer.ShowUsers).Methods("GET")
	userRoute.HandleFunc("/users/{id}", userServer.UpdateUser).Methods("PUT")
	userRoute.HandleFunc("/users/{id}", userServer.UpdateUser).Methods("PUT")
	userRoute.HandleFunc("/users/{id}", userServer.DeleteUser).Methods("DELETE")

	userRoute.Use(middleware.AuthMiddleware)
}
