package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ticker "github.com/tdutanton/concurrent_execution_go/internal/03_ticker"
)

func main() {
	k := flag.Int("k", 1, "Tick interval (seconds)")
	flag.Parse()
	if *k <= 0 {
		fmt.Println("Argument must be positive. Please run again")
		os.Exit(1)
	}
	done := make(chan struct{})
	ticker.RunTicker(done, *k)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	close(done)
	fmt.Println("Termination")
}
