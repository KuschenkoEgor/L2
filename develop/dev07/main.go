package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup

	for _, c := range channels {
		start := time.Now()
		wg.Add(1)
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			fmt.Printf("close channel after %v\n", time.Since(start))
			wg.Done()
		}(c)
	}

	wg.Wait()
	close(out)

	return out
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}
