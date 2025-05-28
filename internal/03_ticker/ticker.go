// Package ticker provides a concurrent ticker function that prints periodic messages.
package ticker

import (
	"fmt"
	"time"
)

// RunTicker starts a goroutine that prints tick messages every k seconds.
// The function uses ANSI color codes to alternate the output color for better visibility.
// It continues printing ticks until the done channel is closed.
//
// Parameters:
//   - done: a receive-only channel used to signal the ticker to stop.
//   - k: tick interval in seconds.
//
// To kill main function with sigterm:
//   - run the program
//   - from other cmd - ps aux | grep main
//     and find something like montoya+   79112  ... /main -k=1 - 79112 is process id
//     then run kill -SIGTERM <id> (like 79112)
func RunTicker(done <-chan struct{}, k int) {
	go func() {
		colors := []string{
			"\033[36m", // Cyan
			"\033[37m", // Light Gray
		}
		i, slept := 1, 0
		for {
			select {
			case <-done:
				return
			default:
				color := colors[i%len(colors)]
				time.Sleep(time.Duration(k) * time.Second)
				slept += k
				fmt.Printf("%sTick %d since %d\n", color, i, slept)
				i++
			}
		}
	}()
}
