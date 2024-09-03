// flow
// user => routers => controllers(logic) => db => reside models
package main

import (
	"log"
	"net/http"

	"github.com/ThanyawitNiti/book-store-management-go/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
