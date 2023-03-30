package main

import (
	"fmt"
	"log"
	"net/http"

	"apitools/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.TestConnection).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8181")
	log.Println("Connected to port 8181")
	log.Fatal(http.ListenAndServe(":8181", router))
}
