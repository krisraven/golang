// closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel's receivers

package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// this goroutine repeatedly receives from jobs with j,more:=<-jobs. In this special 2-value form of receive the more value will be false if jibs has been closed and all values in the channel have already been received. We use this to notify on done when we've worked all of the jobs
	go func() {
		for {
			q, more := <-jobs
			if more {
				fmt.Println("received job", q)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
// this sends 3 jobs to the worker using the jobs channel, then closes the channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
// we await the worker using the synchronisation approach we saw earlier
	<-done
}
