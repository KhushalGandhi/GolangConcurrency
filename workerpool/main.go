package main

import (
	"fmt"
	"time"
)

func main() {
	var numjobs = 10

	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)

	for i := 1; i <= 3; i++ {
		go Worker(i, jobs, results)
	}

	for j := 1; j <= numjobs; j++ {
		jobs <- j // yahan pe j value receive kr rha hai jobs mtlb kaam ko ki haaan job aa gyi hai ise worker ko assign krdo ab
	}
	close(jobs)

	for a := 1; a <= numjobs; a++ {
		<-results // yahan pe hm results receive kr rhe hai
	}

	//ye jo last ke 2 concept thode aise hi lge mujhe

}

func Worker(id int, job chan int, work chan int) {
	for j := range job {
		fmt.Println("worker", id, "started job", j)

		time.Sleep(time.Second)

		fmt.Println("workedId", id, "finished job", j)
		work <- j * 2
	}
}
