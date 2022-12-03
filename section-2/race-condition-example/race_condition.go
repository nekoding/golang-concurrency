package main

import (
	"fmt"
	"sync"
)

var name string
var wg sync.WaitGroup

func updateName(n string) {
	defer wg.Done()
	name = n
}

func main() {
	name = "sarah"

	wg.Add(2)
	go updateName("Jacob")
	go updateName("Brian")
	wg.Wait()

	fmt.Println(name)
}
