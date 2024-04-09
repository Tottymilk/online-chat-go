package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (s *Server) mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/html/index.html")
	if err != nil {
		log.Fatal("oops", err)
	}
	tmpl.Execute(w, nil)
}

func (s *Server) aboutPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./web/html/about.html")
	tmpl.Execute(w, nil)
}

func (s *Server) sendMessage(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s\n", r.PostFormValue("user-message"))
	fmt.Print(message)
	//allMessages = append(allMessages, message)

}
