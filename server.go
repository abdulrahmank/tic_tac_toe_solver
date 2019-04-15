package main

import (
	"github.com/abdulrahmank/solver/tic_tac_toe/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/play", handler.Play)
	log.Panic(http.ListenAndServe(":80", nil))
}
