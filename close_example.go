package main

import "fmt"

func main() {
	ch := make(chan int)

	// Sender goroutine: sends data then closes channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // Signal that no more values will be sent
	}()

	// Receiver using for-range: terminates loop when channel is closed
	for v := range ch {
		fmt.Println("Received:", v)
	}

	fmt.Println("All values received, channel is closed")
}
