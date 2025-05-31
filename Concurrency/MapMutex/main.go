package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map //Map from sync package is used where multiple go routines can write or read from it at same time
	m.Store("apple", 0)
	//When multiple goroutines tries to write to same resources values doesn't get update correctly, so locks are introduced
	var lock sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()         // lock the map resource so that while one goroutines write or read other are blocked
			defer lock.Unlock() //release the resource
			defer wg.Done()
			incrementMap(&m, "apple")
		}()
	}
	wg.Wait()
	val, _ := m.Load("apple")
	fmt.Println("Apples", val)

}

func incrementMap(m *sync.Map, key string) {
	val, _ := m.Load(key)
	intVal := val.(int)
	newVal := intVal + 1
	m.Store(key, newVal)
}
