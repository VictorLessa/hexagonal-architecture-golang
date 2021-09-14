package routes

import (
	"victorLessa/server/application/repositories"
	"victorLessa/server/application/usecases"
	"victorLessa/server/framework/server"

	mux "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


type AuthenticateRoute struct {
	Router *mux.Router
	Db  *gorm.DB
}

func (route *AuthenticateRoute) AuthenticateRoutes() {
	userRepository := repositories.UserRepositoryDb{Db: route.Db}
	authenticateUseCases := usecases.AuthenticateUseCases{UserRepository: userRepository}
	authenticateServer := server.AuthenticateServer{AuthenticateUseCases: authenticateUseCases}

	route.Router.HandleFunc("/signIn", authenticateServer.SignIn).Queries("username", "{username}", "password", "{password}").Methods("GET")
}