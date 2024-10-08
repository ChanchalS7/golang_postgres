package router

import (
	"github.com/ChanchalS7/golang_postgres/middleware"
	"github.com/gorilla/mux"
)

//Router is exported and used in main.go

func Router() *mux.Router {
	router:= mux.NewRouter()

	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock",middleware.GetALLStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newstock",middleware.CreateStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{id}",middleware.DeleteStock).Methods("DELETE","OPTIONS")

	return router
}