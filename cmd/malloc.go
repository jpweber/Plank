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
	"log"
	"time"

	"github.com/jpweber/Plank/pkg/memory"
	"github.com/spf13/cobra"
)

var interval int
var duration int

// mallocCmd represents the malloc command
var mallocCmd = &cobra.Command{
	Use:   "malloc",
	Short: "Allocates memory and hangs on to it. ",
	Long: `Allocates memory growing over time. Can be used in any instance where you 
	need to use up some memory to test things. `,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		buffMap := make(map[int][10 * 1024 * 1024]string)
		go func() {
			x := 0
			for t := range ticker.C {
				fmt.Println("Tick at", t)
				buffMap[x] = memory.Fill()
				x++
			}
		}()

		// 30 seconds == 6 gigs
		// 20 3.5 to == 4 gigs
		// 10 == 1.7 gigs
		time.Sleep(10 * time.Second)
		ticker.Stop()
		// var buffer [100 * 1024 * 1024]string
		// buffMap =
		log.Printf("Done allocating memory.")
		time.Sleep(60 * time.Second)
	},
}

func init() {
	rootCmd.AddCommand(mallocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mallocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mallocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// fill interval. This is how long the ticks are when the fill function is called
	mallocCmd.Flags().IntVarP(&interval, "interval", "i", 1, "Time intervals that memmory will get filled")
	mallocCmd.Flags().IntVarP(&duration, "duration", "d", 3, "How long to run the process for.")
}
