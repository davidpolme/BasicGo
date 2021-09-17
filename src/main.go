package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("hello")
	wg.Add(1)
	go say("World", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)
}
