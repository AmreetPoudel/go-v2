// package main

import (
	"bufio" // For buffered reading and writing
	"fmt"   // For formatted I/O
	"os"    // For file operations
)

func main() {
	// Step 1: Open input file for reading
	inputFile, err := os.Open("app.log") // os.Open returns *os.File that implements io.Reader
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close() // Ensure the file is closed after function ends (good practice)

	// Step 2: Create output file for writing (will overwrite if it already exists)
	outputFile, err := os.Create("output.txt") // os.Create returns *os.File that implements io.Writer
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close() // Close the output file at the end

	// Step 3: Wrap input and output files with buffered reader and writer
	scanner := bufio.NewScanner(inputFile) // Scanner reads line-by-line from inputFile
	writer := bufio.NewWriter(outputFile)  // Writer buffers output to outputFile

	// Step 4: Read each line, add prefix, and write to output file
	for scanner.Scan() { // Scan returns false when EOF is reached
		line := scanner.Text() // Get the current line as string

		// Write the prefixed line to output file
		_, err := fmt.Fprintln(writer, "[INFO]", line) // Writes "[INFO] line content" followed by newline
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}

	// Step 5: Flush buffered data to disk (very important or data may remain in memory)
	writer.Flush()

	fmt.Println("Done. Prefixed lines written to output.txt")
}
