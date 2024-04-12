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
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
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
