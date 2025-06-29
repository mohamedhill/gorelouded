package main

import (
	"fmt"
	"os"
	"strings"

	"goreloaded"
)

func main() {
	filenames := os.Args
	if len(filenames) != 3 {
		fmt.Println("program inputfile:", filenames[1], "||", "outputfile: ??")
		return
	}
	if !strings.HasSuffix(os.Args[1], ".txt") || !strings.HasSuffix(os.Args[2], ".txt") {
		fmt.Println("The input or output file must have a .txt extension.", filenames[1:])
		return
	}
	data, err := os.ReadFile(filenames[1])
	if err != nil {
		fmt.Println("Error: cannot read input file:", err)
		return
	}
	clean := goreloaded.CleanStr(string(data))
	zrox := goreloaded.Gorseloaded(clean)
	err = goreloaded.WriteOutput(filenames[2], zrox)
	if err != nil {
		fmt.Println("Error: cannot write the output file:", err)
		return
	}
}
