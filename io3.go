package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Step 1: Open the source file
	srcFile, err := os.Open("input.txt") // Implements io.Reader
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer srcFile.Close()

	// Step 2: Create the destination file
	destFile, err := os.Create("partial_output.txt") // Implements io.Writer
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer destFile.Close()

	// Step 3: Wrap srcFile in a LimitReader to only allow 100 bytes to be read
	limitedReader := io.LimitReader(srcFile, 100) // Only reads first 100 bytes

	// Step 4: Copy limited data to destination
	bytesCopied, err := io.Copy(destFile, limitedReader)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}

	fmt.Printf("Copied only the first %d bytes from input.txt to partial_output.txt\n", bytesCopied)
}
