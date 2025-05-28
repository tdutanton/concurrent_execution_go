package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	asyncstopwatch "github.com/tdutanton/concurrent_execution_go/internal/01_asyncstopwatch"
)

func main() {
	n := flag.Int("n", 10, "Goroutines count")
	m := flag.Int("m", 1000, "Time to sleep evety of goroutines (ms)")
	flag.Parse()
	if *n < int(asyncstopwatch.MinTime) || *m < int(asyncstopwatch.MinTime) {
		fmt.Println("Arguments must be positive. Please run again")
		os.Exit(1)
	}

	result := asyncstopwatch.GoroutineInfoSlice{}
	var mu sync.Mutex
	wg := &sync.WaitGroup{}
	for i := range *n {
		wg.Add(1)
		go func(i asyncstopwatch.Index) {
			tmp := asyncstopwatch.TakeANap(wg, i, asyncstopwatch.SleepMS(*m))
			mu.Lock()
			result = append(result, tmp)
			mu.Unlock()
		}(asyncstopwatch.Index(i + 1))
	}
	wg.Wait()
	result.SortBySleepTime()
	result.PrintSortedSlice(os.Stdout, *n, *m)
}
