package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./views"))

	http.Handle("/", fs)

	log.Print("Listening on port :3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
