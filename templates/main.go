package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/dcwk/learngo/filereader"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures, err := filereader.ReadStrings("./templates/signatures.txt")
	check(err)
	fmt.Printf("%v\n", signatures)
	html, err := template.ParseFiles("./templates/view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8081", nil)
	log.Fatal(err)
}
