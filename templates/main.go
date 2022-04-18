package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/dcwk/learngo/filereader"
	"github.com/dcwk/learngo/templates/guestbook"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures, err := filereader.ReadStrings("./templates/signatures.txt")
	check(err)

	html, err := template.ParseFiles("./templates/view.html")
	check(err)

	guestbook := guestbook.Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8081", nil)
	log.Fatal(err)
}
