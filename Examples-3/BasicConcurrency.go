package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func tryGet(url string) {
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

func doSum() {
	a, c := 15, 3.5
	d := float64(a) / c
	x := 15 / 3.5
	fmt.Println(x)
	fmt.Println(d)
	b := 2 + 5
	fmt.Println("Hello ", b)
}

func main() {
	defer tryGet("https://facebook.com")
	go tryGet("https://www.github.com")
	go tryGet("https://google.com")
	go tryGet("https://stackoverflow.com")
	go doSum()

	time.Sleep(10 * time.Second)
}
