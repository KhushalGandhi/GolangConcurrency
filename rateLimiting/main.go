package rateLimiting

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5) // buffered

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(20 * time.Millisecond) // see this tick is more of like a channel

	for req := range requests {
		<-limiter //

		// see in above limiter the variable name is written

		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3) // ye time hai

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t // didn't get this part agr 2 jgh se values bhej rha hai
		}
	}()

	burstyRequests := make(chan int, 5) // ye rwquests hao
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
		//ye bhi send kr rha hai
	}
	close(burstyRequests) // ab iska bhi logic thoda clear krna hai ki close pehle kr skte hai ya nhi
	for req := range burstyRequests {
		<-burstyLimiter // ye rokne ke kaam aayega ki jgh hai ya nhi
		fmt.Println("request", req, time.Now())
	}

	 basic sa hai ye hmnein sath sath ek ticker chla diya taaki thoda slow ho jaaye

	mtlb jb values receive ho rhi hai tb hmnein sath mein us limiter se time ka concept receive krna shuru kr diya

	to mainly hm dekhein baaki sb kuch simple hai sirf bustylimiter hi hai jo ki mainly rate limiting ka conept pura
	kr paa rha hai

}
