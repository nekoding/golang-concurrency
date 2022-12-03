package main

import (
	"fmt"
	"sync"
)

var name string
var wg sync.WaitGroup

func updateName(n string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock() // data dikunci
	name = n
	m.Unlock() // data dibuka
}

func main() {

	var Mutex sync.Mutex

	name = "sarah"

	wg.Add(2)
	go updateName("Jacob", &Mutex)
	go updateName("Brian", &Mutex)
	wg.Wait()

	fmt.Println(name)
}
