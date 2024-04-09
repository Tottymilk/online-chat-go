package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/sendMessage", sendMessage).Methods("post")
	log.Fatal(http.ListenAndServe(":8085", router))
}
