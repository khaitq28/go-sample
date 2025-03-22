package main

/*
Goroutines and Channels Tutorial
------------------------------
This tutorial covers concurrency in Go:
1. Goroutines
   - Basic goroutines
   - Goroutine synchronization
   - Goroutine lifecycle
2. Channels
   - Channel basics
   - Buffered channels
   - Channel direction
3. Advanced Concurrency
   - Select statement
   - Worker pools
   - Channel timeouts
4. Best Practices
   - Goroutine leaks prevention
   - Channel patterns
   - Error handling in goroutines
*/

import (
	"fmt"
	"time"
)

// Basic goroutine
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// Goroutine with channel
func sendMessage(ch chan string) {
	ch <- "Hello from channel!"
}

// Goroutine with multiple channels
func processNumbers(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
}

// Goroutine with select
func selectExample(ch1, ch2 chan string) {
	for {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
}

// Goroutine with WaitGroup
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 1. Basic goroutine
	fmt.Println("Basic Goroutine:")
	go sayHello()
	time.Sleep(time.Second)

	// 2. Channel basics
	fmt.Println("\nChannel Basics:")
	ch := make(chan string)
	go sendMessage(ch)
	msg := <-ch
	fmt.Printf("Received: %s\n", msg)

	// 3. Buffered channel
	fmt.Println("\nBuffered Channel:")
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)

	// 4. Channel direction
	fmt.Println("\nChannel Direction:")
	in := make(chan int)
	out := make(chan int)
	go processNumbers(in, out)

	in <- 5
	fmt.Printf("Result: %d\n", <-out)

	// 5. Select statement
	fmt.Println("\nSelect Statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go selectExample(ch1, ch2)

	ch1 <- "Message 1"
	ch2 <- "Message 2"

	// 6. Worker pool
	fmt.Println("\nWorker Pool:")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	// 7. Range over channel
	fmt.Println("\nRange over Channel:")
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for n := range numbers {
		fmt.Printf("Received: %d\n", n)
	}

	// 8. Channel timeout
	fmt.Println("\nChannel Timeout:")
	timeoutCh := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		timeoutCh <- "Message"
	}()

	select {
	case msg := <-timeoutCh:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}
