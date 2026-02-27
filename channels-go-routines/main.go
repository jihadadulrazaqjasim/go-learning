package main

import (
	"fmt"
	"time"
)

type OrderNotification struct {
	OrderID      int
	CustomerName string
	Method       string
}

func notificationWorker(jobs <-chan OrderNotification, done chan<- struct{}) {
	for job := range jobs {
		fmt.Printf("Sending %s notification for Order #%d to %s...\n",
			job.Method, job.OrderID, job.CustomerName)

		time.Sleep(600 * time.Millisecond)

		fmt.Printf("Done: Order #%d\n", job.OrderID)
	}

	done <- struct{}{}
}

func main() {
	jobs := make(chan OrderNotification)
	done := make(chan struct{})

	go notificationWorker(jobs, done)

	for i := 1; i <= 10; i++ {
		jobs <- OrderNotification{
			OrderID:      i,
			CustomerName: fmt.Sprintf("Customer-%d", i),
			Method:       "email",
		}
	}

	close(jobs)

	<-done
	fmt.Println("All notifications processed.")
}
