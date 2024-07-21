package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)

	wg := sync.WaitGroup{}

	b := []int{2, 5, 3, 4}

	for i := 0; i < len(b); i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			Process(val, a)
		}(b[i])

	}

	go func() {
		wg.Wait()

		close(a)
	}()
	for k := range a {
		fmt.Println(k)
	}
}

func Process(u int, ch chan int) {
	ch <- u * 10
}
