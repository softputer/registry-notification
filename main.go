package main

import (
	"github.com/gorilla/mux"
	"log"
	"github.com/softputer/registry-notification/notification"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/events/", notification.EventHandler)
	r.Methods("POST")
	
	log.Fatal(http.ListenAndServe(":9000", r))
}
