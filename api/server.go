package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type Server struct {
	listenAdress string
	conns        map[*websocket.Conn]bool
}

func NewServer(listenAdress string) *Server {
	return &Server{
		listenAdress: listenAdress,
		conns:        make(map[*websocket.Conn]bool),
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("message had been sent"))
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/", s.mainPage)
	router.HandleFunc("/about", s.aboutPage)
	router.HandleFunc("/sendMessage", s.sendMessage).Methods("post")
	router.Handle("/ws", websocket.Handler(s.HandleWS))
	return (http.ListenAndServe(s.listenAdress, router))
}
