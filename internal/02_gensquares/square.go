// Package gensquares provides functions for generating
// integers in a range and computing their squares concurrently.
package gensquares

// Generator returns a read-only channel that emits integers from k to n inclusive.
// If k is greater than n, the values are generated in ascending order from the smaller to the larger.
// The values are sent from a separate goroutine, and the channel is closed after the sequence is sent.
func Generator(k, n int) <-chan int {
	c := make(chan int)
	if k > n {
		k, n = n, k
	}
	go func() {
		defer close(c)

		for i := k; i <= n; i++ {
			c <- i
		}
	}()
	return c
}

// Square returns a read-only channel that emits the squares of integers
// received from the input read-only channel c.
// It processes the input values in a separate goroutine and closes the output
// channel after all values have been squared and sent.
func Square(c <-chan int) <-chan int {
	s := make(chan int)
	go func() {
		defer close(s)

		for v := range c {
			s <- v * v
		}
	}()
	return s
}
