package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func tryGet(url string) {

	// Defer the call to waitgroup's done to tell that goroutine is finished
	defer wg.Done()

	fmt.Println("Firing Get for ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Defering response body closure for ", url)
	defer response.Body.Close()

	fmt.Println("Reading response body for ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response size for url ", url, " is ", len(body))
}

func main() {
	wg.Add(10)
	go tryGet("https://www.github.com")
	go tryGet("https://google.com")
	go tryGet("https://stackoverflow.com")
	go tryGet("https://facebook.com")
	go tryGet("https://oracle.com")
	go tryGet("https://go.org")
	go tryGet("https://cisco.edu")
	go tryGet("https://coursera.org")
	go tryGet("https://security.us.ces.cisco.com")

	// Wait for goroutines to finish
	wg.Wait()

	fmt.Println("Work is over, holidays are here!!!")
}
