package main

import (
	"fmt"
	"sync"
)

var counter = 0
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	m.RLock()
	fmt.Printf("Hello: #%d\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
