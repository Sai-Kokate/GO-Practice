package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine is done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second * 5)
	fmt.Printf("Worker %d done\n", id)
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add 1 to the counter for each worker
		go worker(i, &wg)
	}

	counter := Counter{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Wait until all workers are done
	fmt.Println("All workers have finished")
	fmt.Println("Final value:", counter.GetValue())
}
