package main

import (
	"encoding/json"
	"net/http"
	"os"

	"victorLessa/server/framework/cmd/routes"
	"victorLessa/server/framework/database"

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
	
	userRoute := routes.UserRoute{Router: a.Router, Db: a.Db }
	authenticateRoute := routes.AuthenticateRoute{Router: a.Router, Db: a.Db }

	userRoute.UserRoutes()
	authenticateRoute.AuthenticateRoutes()

}

func (a *App) Run() {
	http.ListenAndServe(":8010", a.Router)
}

func main() {
	a := App{}

	a.Initialize()
	a.Run()
	
}