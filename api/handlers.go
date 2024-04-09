package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func (s *Server) mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("index.html").ParseFiles("index.html")
	tmpl.Execute(w, nil)
}

func (s *Server) sendMessage(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s\n", r.PostFormValue("user-message"))
	fmt.Print(message)
	//allMessages = append(allMessages, message)

}
