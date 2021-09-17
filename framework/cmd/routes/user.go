package routes

import (
	"net/http"
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
	userRoute.HandleFunc("/users", middleware.Adapt(http.HandlerFunc(userServer.CreateUser), middleware.AuthMiddleware).ServeHTTP).Methods("POST")
	userRoute.HandleFunc("/users", middleware.Adapt(http.HandlerFunc(userServer.IndexUsers), middleware.AuthMiddleware).ServeHTTP).Methods("GET")
	userRoute.HandleFunc("/users/{id}", middleware.Adapt(http.HandlerFunc(userServer.ShowUsers), middleware.AuthMiddleware).ServeHTTP).Methods("GET")
	userRoute.HandleFunc("/users/{id}", middleware.Adapt(http.HandlerFunc(userServer.UpdateUser), middleware.AuthMiddleware).ServeHTTP).Methods("PUT")
	userRoute.HandleFunc("/users/{id}", middleware.Adapt(http.HandlerFunc(userServer.DeleteUser), middleware.AuthMiddleware).ServeHTTP).Methods("DELETE")

}
