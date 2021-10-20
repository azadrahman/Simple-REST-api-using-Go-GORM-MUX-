package main

import (
	"bookmanagement/src/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Println("server running.....")
	log.Fatal(http.ListenAndServe(":9000", r))
}
