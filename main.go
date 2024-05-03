package main

import (
	"flag"
	"log"
	"main/api"
)

func main() {
	listenAddr := flag.String("p", ":8085", "server adress")
	s := api.NewServer(*listenAddr)
	log.Fatal(s.Start())
}
