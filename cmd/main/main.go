package main

import (
	"log"
	"net/http"

	"github.com/rudraprasaaad/BMS/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
