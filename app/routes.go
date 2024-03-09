package app

import "github.com/IdaDanuartha/online-shop-golang/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}