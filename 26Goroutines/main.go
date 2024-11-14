package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex // pointer
var signals = []string{"test"}

func main() {

	websitelist := []string{"https://ineuron.ai", "https://blog.learncodeonline.in", "https://nirajsalunke.netlify.app", "https://go.dev/"}

	for _, val := range websitelist {
		go GetStatusCode(val)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func GetStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d Status code for %s\n", result.StatusCode, endpoint)
}
