package disk

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func Write(size int64) bool {
	t := time.Now().UTC()

	f, err := os.OpenFile("output", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for i := int64(0); i < size; i++ {
		n, err := io.WriteString(f, "0")
		// n, err := f.Write([]byte{1})
		if err != nil {
			fmt.Println("Error writing to file", n, err)
		}
	}

	err = f.Close()
	//end if disk operations
	fmt.Println("Time to write file", time.Since(t))

	if err != nil {
		log.Fatal("Failed to close file")
		return false
	}

	return true
}

func Read(fd *os.File) {
	file, err := os.OpenFile("output", os.O_RDONLY, 666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t := time.Now().UTC()

	b, err := ioutil.ReadAll(file)
	fmt.Println("Time to Read file", time.Since(t))
	fmt.Print(len(b))

	return

}

// func main() {

// 	fileSize := flag.Int("s", 100, "Size of file to write in MB")
// 	// Once all flags are declared, call `flag.Parse()`
// 	// to execute the command-line parsing.
// 	flag.Parse()

// 	size := int64(*fileSize * 1024 * 1024)
// 	_ = os.Remove("output")

// 	fd, err := os.Create("output")
// 	if err != nil {
// 		log.Fatal("Failed to create output")
// 	}
// 	fd.Close()

// 	// start of disk write
// 	if write(size) {
// 		// if write succeeds then to the read
// 		fd, _ = os.Open("output")
// 		read(fd)
// 	}

// }
