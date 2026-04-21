package routes

import (
	"librarymanagement-system/handlers"

	"github.com/gorilla/mux"
)

func RagisterRoute(r *mux.Router){
	r.HandleFunc("/create",handlers.CreateBook).Methods("POST")
	r.HandleFunc("/",handlers.GetAll).Methods("GET")
	r.HandleFunc("/{id}",handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/{id}",handlers.DeleteById).Methods("DELETE")
	r.HandleFunc("/data/delete-all",handlers.DeleteAll).Methods("POST")
}