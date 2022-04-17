package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func formattedViewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("<h1>Hello web!</h1>")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/formatted-hello", formattedViewHandler)

	err := http.ListenAndServe("localhost:8081", nil)
	log.Fatal(err)
}
