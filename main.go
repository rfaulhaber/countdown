package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("countdown: received no arguments")
		os.Exit(1)
	}

	dur, err := time.ParseDuration(os.Args[1])

	if err != nil {
		fmt.Println("countdown: received invalid time:", os.Args[1])
		os.Exit(1)
	}

	start := time.Now()

	c := time.After(dur)

	for {
		select {
		case <-c:
			fmt.Println("\ndone!")
			os.Exit(0)
		default:
			fmt.Print("\033[K", dur - time.Since(start).Round(time.Second), "\r")
		}
	}
}
