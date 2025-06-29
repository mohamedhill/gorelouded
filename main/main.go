package main

import (
	"fmt"
	"os"

	"goreloaded"
)

func main() {
	filenames := os.Args
	if len(filenames) < 3 {
		fmt.Println("program inputfile:",filenames[1],"||", "outputfile: ??")
		return
	}
	if !goreloaded.ProtectedFile(filenames[0], filenames[1]) {
		fmt.Println("This program only works with 'sample.txt' and 'result.txt'not:", filenames[1:])
		return
	}
	data, err := os.ReadFile(filenames[1])
	if err != nil {
		fmt.Println("program can't read  input file>>:", err)
		return
	}
	clean := goreloaded.CleanStr(string(data))
	zrox := goreloaded.Gorseloaded(clean)
	err = goreloaded.WriteOutput(filenames[2], zrox)
	if err != nil {
		fmt.Println("program can't write output file :", err)
		return
	}
}
