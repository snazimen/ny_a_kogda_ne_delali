package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8888"
	}

	webDir := "./web"
	r := chi.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(webDir)))

	serverAdress := fmt.Sprintf("localhost:%s", port)
	log.Println("listen on " + serverAdress)
	if err = http.ListenAndServe(serverAdress, http.FileServer(http.Dir(webDir))); err != nil {
		log.Panicf("start server error:%s", err.Error())
	}
}
