package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var allMessages []string

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("index.html").ParseFiles("index.html")
	tmpl.Execute(w, allMessages)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s\n", r.PostFormValue("user-message"))
	fmt.Print(message)
	allMessages = append(allMessages, message)

}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/sendMessage", sendMessage).Methods("post")
	log.Fatal(http.ListenAndServe(":8085", router))
}

func main() {
	allMessages = append(allMessages, "wassup beijing\n")
	handleRequests()
}
