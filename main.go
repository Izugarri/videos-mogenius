package main

import (
	"log"
	"os"
	"fmt"
	"net/http"
)


func main() {
	fmt.Println("starting your http server!")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}
