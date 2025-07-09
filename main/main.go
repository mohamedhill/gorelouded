package main

import (
	"fmt"
	"os"
	"strings"

	"goreloaded/textproc" // Replace with the actual module name if different
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input. Please provide exactly two arguments: input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if !strings.HasSuffix(inputFile, ".txt") || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: Both input and output files must have a .txt extension.")
		return
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error: cannot read input file: %v\n", err)
		return
	}

	input := string(data)
	processed := textproc.Gorseloaded(input)

	if err := textproc.WriteOutput(outputFile, processed); err != nil {
		fmt.Printf("Error: cannot write output file: %v\n", err)
		return
	}
}
