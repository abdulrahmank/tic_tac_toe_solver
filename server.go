package main

import (
	"fmt"
	"github.com/abdulrahmank/solver/tic_tac_toe/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	http.HandleFunc("/play", handler.Play)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
