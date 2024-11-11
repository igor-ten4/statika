package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	listen := os.Getenv("STATIKA_LISTEN")
	if listen == "" {
		listen = ":8080"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Printf("Listening on %s", listen)
	err := http.ListenAndServe(listen, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
