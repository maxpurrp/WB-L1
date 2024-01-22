package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int        // value of counter
	mu    sync.Mutex // mutex for savety acces
}

// increment increases the counter value by 1
func (c *Counter) Increment() {
	c.mu.Lock()         //lock mutex before entering the critical section
	defer c.mu.Unlock() //releasing the mutex when exiting the function
	c.value++
}

// return current value of count
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	//create an instance of the counter structure
	counter := &Counter{}

	// use waitGroup for wait for the completion all goroutines
	var wg sync.WaitGroup

	// run several goroutines to increment value in parallel
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	//waiting all goroutines
	wg.Wait()

	fmt.Println("Total count:", counter.GetValue())
}
