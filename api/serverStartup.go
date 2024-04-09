package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	s := server.Server
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/sendMessage", sendMessage).Methods("post")
	log.Fatal(http.ListenAndServe(s.listenAdress, router))
}
