package memory

import "log"

func Fib(n uint64) uint64 {
	if n <= 2 { // terminate recursion
		return 1
	}
	// create a channel for subtask results
	res := make(chan uint64, 2)
	// fork 2 subtasks, they will send results to the res channel
	go func() {
		res <- Fib(n - 1)
	}()
	go func() {
		res <- Fib(n - 2)
	}()
	return <-res + <-res
}

func Avg(values [3]float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}

func Debug(message interface{}) {
	// if verbose {
	log.Println(message)
	// }
}
