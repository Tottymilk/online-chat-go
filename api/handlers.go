package api

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

func (s *Server) mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/html/index.html")
	tmpl.Execute(w, nil)
}

func (s *Server) aboutPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/html/about.html")
	tmpl.Execute(w, nil)
}

func (s *Server) sendMessage(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s\n", r.PostFormValue("user-message"))
	fmt.Print(message)

}

func (s *Server) HandleWS(ws *websocket.Conn) {
	var mutex sync.Mutex
	fmt.Println("new conn from client:", ws.RemoteAddr())

	mutex.Lock()
	s.conns[ws] = true
	mutex.Unlock()

	s.readLoop(ws)
}
