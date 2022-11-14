package main

import (
	"log"
	"net/http"

	"github.com/MohamedSawahZC/book_management_system/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



