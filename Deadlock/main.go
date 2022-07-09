package main

import "sync"

var wg sync.WaitGroup

// Channels are a way to communcate between goroutines.
func doStuff(c1 chan int, c2 chan int) {
	<-c1    // waits to recieve from c1
	c2 <- 1 // writes to c2
	wg.Done()
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	wg.Add(2)
	go doStuff(c1, c2)
	go doStuff(c2, c1)
	wg.Wait()
}
