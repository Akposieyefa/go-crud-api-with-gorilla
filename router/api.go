package router

import (
	"log"
	"net/http"

	"github.com/akposieyefa/crud-gorilla/initializers"
	"github.com/gorilla/mux"
)

func LoadRouters() {

	r := mux.NewRouter()

	r.HandleFunc("/users", initializers.GetUsers).Methods("GET")
	r.HandleFunc("/users", initializers.CreateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", initializers.GetSingleUsers).Methods("GET")
	r.HandleFunc("/users/{id}", initializers.UpdateUsers).Methods("PATCH")
	r.HandleFunc("/users/{id}", initializers.DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:9000", r))

}
