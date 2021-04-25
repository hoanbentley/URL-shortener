package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/hoanbentley/URL-shortener/internal/controller"
	"net/http"
)

var db string

func main() {
	flag.StringVar(&db, "db", "", "db run")
	flag.Parse()
	router := mux.NewRouter()
	todo := controller.NewToDoService()
	router.HandleFunc("/login", todo.GetAuthToken).Methods(http.MethodGet)
	router.HandleFunc("/admin/list", todo.ListUrl).Methods(http.MethodGet)
	router.HandleFunc("/admin/search", todo.SearchUrl).Methods(http.MethodPost)
	router.HandleFunc("/admin/delete/{id}", todo.DeleteUrl).Methods(http.MethodDelete)
	router.HandleFunc("/create", todo.CreateUrl).Methods(http.MethodPost)
	router.HandleFunc("/{id}", todo.RedirectUrl).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}
