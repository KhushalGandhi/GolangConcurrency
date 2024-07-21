package main

import "fmt"

func main() {
	a := make(chan int)

	b := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(b); i++ {
		go process(b[i], a)
	}

	for i := 0; i < len(b); i++ {
		k := <-a
		fmt.Println(k)
	}

}

func process(c int, ch chan int) {
	ch <- c * 10
}
