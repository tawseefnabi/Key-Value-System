package main

import (
	controller "KeyValueMain/Controller"
	repository "KeyValueMain/Repository"
	service "KeyValueMain/Service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	addr string = "127.0.0.1:"
	port string = "8008"
)

func main() {

	fmt.Println("==================== Init Controller ====================")
	fmt.Println("==================== Repository init ====================")
	repository := repository.NewRepository("../keyValue.db")
	service := service.NewService(repository)
	fmt.Println("service init", service)
	controller := controller.NewController(service)
	fmt.Println("controller init", controller)
	fmt.Println("===============================================================")
	router := mux.NewRouter()
	router.HandleFunc("/api/put", controller.put).Methods("PUT")

	srv := http.Server{
		Handler: router,
		Addr:    addr + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running at: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
