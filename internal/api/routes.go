package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mustafaerbay/cleango/internal/api/handlers"
)

/*
This example shows a simple implementation of the route setup for handling CRUD operations on users. 
The App struct contains a Router object of type mux.Router and a UserHandlers object of type handlers.
UserHandler. The SetupRoutes method sets up the various routes that correspond to the different HTTP methods.

Note that this example assumes that the handlers.UserHandler object has already been defined in 
the internal/api/handlers/user.go file. The App simply acts as a coordinator for the HTTP routes and
the handlers that correspond to those routes. In a larger application, there may be more route setup 
methods and more handlers that correspond to different parts of the application.
*/
type App struct {
	Router       *mux.Router
	UserHandlers *handlers.UserHandler
}

func (a *App) SetupRoutes() {
	a.Router.HandleFunc("/users", a.UserHandlers.List).Methods(http.MethodGet)
	a.Router.HandleFunc("/users", a.UserHandlers.Create).Methods(http.MethodPost)
	a.Router.HandleFunc("/users/{id}", a.UserHandlers.Get).Methods(http.MethodGet)
	a.Router.HandleFunc("/users/{id}", a.UserHandlers.Update).Methods(http.MethodPut)
	a.Router.HandleFunc("/users/{id}", a.UserHandlers.Delete).Methods(http.MethodDelete)
}
