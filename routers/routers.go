package routers

import (
	"crud-api-mysql/controller"
	"crud-api-mysql/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouters() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc("/", controller.HelloWeb).Methods(http.MethodGet)
	router.HandleFunc("/create", controller.Create).Methods(http.MethodPost)
	router.HandleFunc("/delete/{id}", controller.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/list", controller.ListAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", controller.UpdateUser).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":5000", router))
}
