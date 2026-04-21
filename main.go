package main

import (
	"fmt"
	"librarymanagement-system/config"
	"librarymanagement-system/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.CoonectDb()

	r:=mux.NewRouter()

	routes.RagisterRoute(r)
	fmt.Println(("Server has been started"))
	http.ListenAndServe(":8080",r)
}