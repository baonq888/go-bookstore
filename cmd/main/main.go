package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("Server running on port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010",r))
	
}