package main

import (
	"fmt"
	"sync"
)

// func doStuff() {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		for j := 0; j < 100000; j++ {
// 		}
// 		fmt.Println(i)
// 	}
// }

var wg sync.WaitGroup
var on sync.Once

func setup() {
	fmt.Println("Init")
}

func doStuff() {
	on.Do(setup) // executes just once no matter how many goroutines are created
	fmt.Println("Hello, World!")
	wg.Done()
}

func main() {
	wg.Add(2)
	go doStuff()
	go doStuff()
	wg.Wait()
}
