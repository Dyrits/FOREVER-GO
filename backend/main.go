package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "The API is up and running!")
		if err != nil {
			log.Fatal(fmt.Sprintf("An error occured while writing to the response: %v", err))
		}
	})

	err := http.ListenAndServe(":4500", nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("An error occured while starting the server: %v", err))
	}
}
