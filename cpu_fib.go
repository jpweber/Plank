package main

import (
	"flag"
	"log"
	"time"
)

func fib(n uint64) uint64 {
	if n <= 2 { // terminate recursion
		return 1
	}
	// create a channel for subtask results
	res := make(chan uint64, 2)
	// fork 2 subtasks, they will send results to the res channel
	go func() {
		res <- fib(n - 1)
	}()
	go func() {
		res <- fib(n - 2)
	}()
	return <-res + <-res
}

func main() {

	iter := flag.Int("i", 30, "number of iterations to run fib sequence. Default of 30")
	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	t := time.Now().UTC()
	fib(uint64(*iter))
	log.Println("Time to run with", *iter, "iterations", time.Since(t))
}
