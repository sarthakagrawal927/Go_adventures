package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"google", "facebook", "amazon", "stackoverflow", "github", "instagram", "flipkart", "snapdeal", "gitlab"}

	c := make(chan string) // declaring channel

	for _, link := range links {
		go checkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l) //Now l can change as the function will be getting its copy
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get("http://" + link + ".com")
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
