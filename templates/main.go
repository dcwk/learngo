package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/dcwk/learngo/filereader"
	"github.com/dcwk/learngo/templates/guestbook"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures, err := filereader.ReadStrings("./templates/signatures.txt")
	check(err)

	html, err := template.ParseFiles("./templates/html/view.html")
	check(err)

	guestbook := guestbook.Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func addHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("./templates/html/add.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)
}

func saveHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	_, err := writer.Write([]byte(signature))
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/add", addHandler)
	http.HandleFunc("/guestbook/save", saveHandler)
	err := http.ListenAndServe("localhost:8081", nil)
	log.Fatal(err)
}
