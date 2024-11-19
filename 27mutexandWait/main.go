package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condtion and go routines")
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	// iif Immediately Involed Functions
	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("1 routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("2 routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3 routine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
