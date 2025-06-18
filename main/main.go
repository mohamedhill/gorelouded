package main

import (
	"fmt"
	"os"
)

func main() {
	filenames := os.Args
	data, err := os.ReadFile(filenames[1])

	if err != nil {
		fmt.Println("error", err)
	} else {
		erre := os.WriteFile(filenames[2], []byte(data), 0644)
		if erre != nil {
			fmt.Println("error", erre)
		}
	}
}
