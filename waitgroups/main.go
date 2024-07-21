package main

import (
	"fmt"
	"sync"
)

func main() {
	//a := make(chan int)

	g := []int{2, 6, 3, 5}

	wg := sync.WaitGroup{}

	for i := 0; i < len(g); i++ {
		wg.Add(1)
		go Process(g[i], &wg)
	}
	wg.Wait()
	fmt.Println("done")

}

func Process(r int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(r * 10)
}
