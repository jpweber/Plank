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
	"log"
	"os"

	"github.com/jpweber/Plank/pkg/disk"
	"github.com/spf13/cobra"
)

var fileSize int

// diskwriteCmd represents the diskwrite command
var diskwriteCmd = &cobra.Command{
	Use:   "diskwrite",
	Short: "Writes/Reads file to disk",
	Long:  `Test Write and Read speeds and creates disk i/o`,
	Run: func(cmd *cobra.Command, args []string) {
		// fileSize := flag.Int("s", 100, "Size of file to write in MB")
		// // Once all flags are declared, call `flag.Parse()`
		// // to execute the command-line parsing.
		// flag.Parse()

		size := int64(fileSize * 1024 * 1024)
		_ = os.Remove("output")

		fd, err := os.Create("output")
		if err != nil {
			log.Fatal("Failed to create output")
		}
		fd.Close()

		// start of disk write
		if disk.Write(size) {
			// if write succeeds then to the read
			fd, _ = os.Open("output")
			disk.Read(fd)
		}
	},
}

func init() {
	rootCmd.AddCommand(diskwriteCmd)

	diskwriteCmd.Flags().IntVarP(&fileSize, "size", "s", 100, "Size of file to write in MB")
}
