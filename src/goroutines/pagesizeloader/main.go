package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Url  string
	Size int
}

func main() {
	myChannel := make(chan Page)
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}

	for _, url := range urls {
		go responseSize(url, myChannel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-myChannel)
	}
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{Url: url, Size: len(body)}
}
