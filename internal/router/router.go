// internal/router/router.go
package router

import (
	"net/http"
	"github.com/rezalaal/coral/internal/handler"
	"github.com/rezalaal/coral/internal/repository/interfaces"
)

func NewRouter(userRepo interfaces.UserRepository) http.Handler {
	mux := http.NewServeMux()

	// Handlers
	userHandler := handler.NewUserHandler(userRepo)
	// restaurantHandler := handler.NewRestaurantHandler(restaurantRepo)

	mux.HandleFunc("/users", userHandler.GetUsers)           // GET
	mux.HandleFunc("/users/create", userHandler.CreateUser)  // POST

	// mux.HandleFunc("/restaurants", restaurantHandler.GetRestaurants)          // GET
	// mux.HandleFunc("/restaurants/create", restaurantHandler.CreateRestaurant) // POST


	return mux
}
