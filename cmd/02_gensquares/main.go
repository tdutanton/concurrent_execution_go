package main

import (
	"flag"
	"fmt"

	gensquares "github.com/tdutanton/concurrent_execution_go/internal/02_gensquares"
)

func main() {
	k := flag.Int("k", 1, "First integer")
	n := flag.Int("n", 9, "Second integer")
	flag.Parse()
	for v := range gensquares.Square(gensquares.Generator(*k, *n)) {
		fmt.Println(v)
	}
}
