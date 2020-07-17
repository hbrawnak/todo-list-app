package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.HealthCheckHandler).Methods("GET", "OPTIONS")

	r := router.PathPrefix("/api").Subrouter()
	r.HandleFunc("/task", middleware.GetAllTask).Methods("GET", "OPTIONS")

	return router
}
