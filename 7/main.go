package main

import (
	"fmt"
	"sync"
)

// creating mutex for secure access
var mutex sync.Mutex

func main() {
	// waitGroup for synchronyse goroutine
	var wg sync.WaitGroup
	//create map for saving data
	myData := make(map[string]int)
	//count goroutine, which competitively save data in map
	countGoroutins := 4
	wg.Add(countGoroutins)
	for i := 1; i < countGoroutins+1; i++ {
		go func(i int) {
			defer wg.Done()
			// block mutex for safety access
			mutex.Lock()
			//deferred unblock for another goroutines
			defer mutex.Unlock()
			//save key and value
			key := fmt.Sprintf("Key%d", i)
			value := i * 10
			myData[key] = value
		}(i)
	}
	//waitgroup to wait for everyone goroutine wrote data
	wg.Wait()

	for key, value := range myData {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
