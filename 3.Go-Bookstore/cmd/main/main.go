package main

import (
	"log"
	"net/http"

	"github.com/eko360k/LearnGoProgrammingByBuildingElevenProjects/3.Go-Bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.ResgisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
