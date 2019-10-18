package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/febrielven/go_mysql_api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe("localhost:3030", r))
}
