// Package asyncstopwatch provides utilities to run concurrent goroutines
// that simulate work by sleeping for a randomized amount of time, and
// collect timing information for analysis or benchmarking.
package asyncstopwatch

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// ICli is the default label used to describe goroutine indices in CLI output.
const ICli string = "Goroutine"

// STCLi is the default label used to describe sleep times (in milliseconds) in CLI output.
const STCLi string = "Sleep time (ms)"

// Index represents the identifier of a goroutine.
type Index int

// SleepMS represents the duration a goroutine sleeps, in milliseconds.
type SleepMS int

// GoroutineInfo stores data about a single goroutine execution,
// including its index and how long it slept.
type GoroutineInfo struct {
	I         Index   // The index (ID) of the goroutine
	SleepTime SleepMS // The sleep duration in milliseconds
}

// GoroutineInfoSlice is a slice of GoroutineInfo structs.
// It can be used to collect and sort data from multiple goroutines.
type GoroutineInfoSlice []GoroutineInfo

// TakeANap launches a goroutine that sleeps for a random duration between 1 ms
// and the provided maximum sleep time (m). It marks the WaitGroup as done
// and returns a GoroutineInfo struct with the goroutine's index and actual sleep time.
func TakeANap(wg *sync.WaitGroup, i Index, m SleepMS) GoroutineInfo {
	defer wg.Done()
	sleep := randTime(m)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	return GoroutineInfo{i, sleep}
}

// SortBySleepTime sorts the GoroutineInfoSlice in descending order by SleepTime.
// Goroutines that slept the longest will appear first.
func (s *GoroutineInfoSlice) SortBySleepTime() {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i].SleepTime > (*s)[j].SleepTime
	})
}

// MinTime is the minimum random sleep duration in milliseconds.
const MinTime SleepMS = 1

// randTime returns a random SleepMS duration between minTime and the provided maximum (m).
func randTime(m SleepMS) SleepMS {
	return SleepMS(rand.Intn(int(m-MinTime+1))) + MinTime
}

// PrintSortedSlice prints each GoroutineInfo element in the slice to the given writer,
// including its index, sleep duration, and a simple ASCII bar chart representing the sleep time.
//
// The bar chart is composed of '#' characters, where the number of characters is
// proportional to the goroutine's sleep duration. The output is meant to provide
// a quick visual comparison of the sleep times.
//
// Assumes the slice is already sorted in descending order by SleepTime.
func (s GoroutineInfoSlice) PrintSortedSlice(writer io.Writer, n int, m int) {
	colors := []string{
		"\033[36m", // Cyan
		"\033[37m", // Light Gray
	}
	step := s[0].SleepTime / 10
	for i, v := range s {
		color := colors[i%len(colors)]
		bar := ""
		for i := int64(0); i < int64(v.SleepTime); i += int64(step) {
			bar += "#"
		}
		fmt.Fprintf(writer, "%s|%-s %*d | %s %*d | %s\n", color, ICli,
			len(fmt.Sprintf("%d", n)), v.I, STCLi, len(fmt.Sprintf("%d", m)), v.SleepTime, bar)
	}
}
