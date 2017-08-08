package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// Info from config file
type Average struct {
	Fib float64
}

// Reads info from averages file
func ReadConfig() Average {
	configfile := "benchmarks.conf"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var average Average
	if _, err := toml.DecodeFile(configfile, &average); err != nil {
		log.Fatal(err)
	}
	return average
}

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

func avg(values [3]float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}

func debug(message interface{}) {
	if *verbose {
		log.Println(message)
	}
}

var verbose *bool

func main() {

	iter := flag.Int("i", 30, "number of iterations to run fib sequence. Default of 30")
	verbose = flag.Bool("verbose", false, "show human readable output")
	flag.Parse()

	averages := ReadConfig()
	times := [3]float64{}
	for i := 0; i < 3; i++ {
		t := time.Now().UTC()
		fib(uint64(*iter))
		times[i] = time.Since(t).Seconds()
		debug(time.Since(t).Seconds())
	}

	result := avg(times)
	if result > averages.Fib+1 {
		debug(fmt.Sprintf("%f %f", result, averages.Fib))
		os.Exit(1)
	}
	debug(fmt.Sprintf("average fib time (seconds): %f user provided threshold (seconds):%f", result, averages.Fib))
	os.Exit(0)

}
