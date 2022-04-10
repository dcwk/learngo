package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	myChannel := make(chan string)
	go responseSize("https://example.com")
	go responseSize("https://golang.org")
	go responseSize("https://golang.org/doc")

	fmt.Println(<-myChannel)
}

func responseSize(url string, channel chan string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- url + " " + string(len(body))
}
