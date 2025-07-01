// package main

import (
	"bufio"   // Provides buffered I/O operations
	"fmt"     // For formatted printing
	"strings" // To create an in-memory Reader from a string
)

func main() {

	// Create an in-memory string Reader from a static string.
	// `strings.NewReader(...)` returns a value that implements `io.Reader`.
	// This simulates a stream of data like reading from a file or network.
	stringReader := strings.NewReader("this is data to practice bufio\n")

	// Wrap the stringReader with a buffered reader.
	// `bufio.NewReader()` adds buffering over any io.Reader.
	// It reads large chunks internally, improving performance and allowing flexible reading methods.
	reader := bufio.NewReader(stringReader)

	// Create a byte slice (buffer) with a capacity of 20 bytes.
	// We use this to store the bytes that we read from the reader.
	// Variable `data` will contain raw byte data read from the source.
	data := make([]byte, 20)

	// Call `reader.Read(data)` to read up to 20 bytes into `data`.
	// It returns:
	// - `n`: number of bytes actually read (could be less than 20)
	// - `err`: error if any occurred during the read (EOF or something else)
	n, err := reader.Read(data)

	// If there was an error while reading (e.g., file closed, I/O issue), handle it.
	if err != nil {
		fmt.Println("we got error: ", err)
		return
	}

	// At this point:
	// - `data[:n]` contains the actual bytes read.
	// - `data[n:]` contains unused zero bytes (padding).
	// `data` is still a byte slice, so we use `%s` to print it as a string.
	fmt.Printf("we read data of %d bytes and data is : %s\n", n, data[:n])

	//reading till deliminator but this will start from where we left to read beause we are using same readere
	abc, _ := reader.ReadString('\n')

	fmt.Println(abc)
}

//so steps are
// Identify your data source
// → This could be a file, a string, standard input, network connection, or HTTP response.

// Get a reader from that source
// → Ensure it implements io.Reader.

// Wrap that reader with bufio.NewReader()
// → Adds buffering and methods for efficient and flexible reading.

// Create a buffer or use a read method
// → Either allocate a []byte slice or use helper methods like ReadString() or Scanner.

// Read data and process it as needed
// → You now read into your buffer, slice, or string — and work with the result (store, print, transform, etc).
