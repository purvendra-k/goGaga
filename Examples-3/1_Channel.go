package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func tryGet(url string, nums chan int) {

	// Defer the call to waitgroup's done to tell that goroutine is finished
	defer wg.Done()
	/*
		if url == "https://www.github.com" {
			time.Sleep(10 * time.Second)
			fmt.Println("Github woke up")
		}
	*/
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

	// Write the response body size in channel
	nums <- len(body)
	fmt.Println("chan len url ", url)
	//fmt.Println("Response size for url ", url, " is ", len(body))
}

func main() {

	// Unbuffered channel
	// If the channel is unbuffered, the sender blocks until the receiver has received the value
	nums := make(chan int)
	wg.Add(3)
	go tryGet("https://www.github.com", nums)
	//go tryGet("https://google.com", nums)
	go tryGet("https://stackoverflow.com", nums)

	fmt.Println("chan len ", len(nums))
	fmt.Println(<-nums)

	fmt.Println("chan len ", len(nums))
	// Wait for goroutines to finish
	wg.Wait()
	fmt.Println("Threads are done")
	close(nums)

	fmt.Println("Work is over, holidays are here!!!")
}
