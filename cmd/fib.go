// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jpweber/Plank/pkg/memory"
	"github.com/spf13/cobra"
)

var iter int

// fibCmd represents the fib command
var fibCmd = &cobra.Command{
	Use:   "fib",
	Short: "Runs fibonnaci sequences to stress memory and cpu",
	Long:  `Will calculate fibonacci out to the specified depth. Stresses memory and Cpu.`,
	Run: func(cmd *cobra.Command, args []string) {

		times := [3]float64{}
		for i := 0; i < 3; i++ {
			t := time.Now().UTC()
			memory.Fib(uint64(iter))
			times[i] = time.Since(t).Seconds()
			memory.Debug(time.Since(t).Seconds())
		}

		result := memory.Avg(times)

		memory.Debug(fmt.Sprintf("average fib time (seconds): %f", result))
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(fibCmd)

	fibCmd.Flags().IntVarP(&iter, "iterations", "i", 30, "number of iterations to run fib sequence. Default of 30")

}
