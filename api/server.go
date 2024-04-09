package api

import (
	"net/http"
)

type Server struct {
	listenAdress string
}

func NewServer(listenAdress string) *Server {
	return &Server{
		listenAdress: listenAdress,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.mainPage)
	http.HandleFunc("/about", s.aboutPage)
	http.HandleFunc("/sendMessage", s.sendMessage)
	return (http.ListenAndServe(s.listenAdress, nil))
}
