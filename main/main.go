package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"goreloaded"
)

func main() {
	filenames := os.Args
	if len(filenames) < 3 {
		fmt.Println("Usage: program inputfile outputfile")
		return
	}

	data, err := os.ReadFile(filenames[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var zrox []string
	clean := goreloaded.CleanStr(string(data))
	zrox = goreloaded.StringToSlice(clean)
	fmt.Println("start:",clean)
	zrox = processTags(zrox)
	fmt.Println("str after flags:", zrox)
	clean = strings.Join(zrox, " ")

	clean = normalizePunctuation(clean)
	fmt.Println("str after ponc:", clean)

	
	zrox = goreloaded.StringToSlice(clean)

	zrox = goreloaded.Cleanslice(zrox)
	fmt.Println("final:",zrox)
	

	err = writeOutput(filenames[2], zrox)
	if err != nil {
		fmt.Println("Error writing output:", err)
	}
}


func normalizePunctuation(input string) string {
	runes := []rune(input)
	var result []rune

	for i := 0; i < len(runes); {
		r := runes[i]

		if goreloaded.Runponc(r) {
			start := i
			for i+1 < len(runes) && goreloaded.Runponc(runes[i+1]) {
				i++
			}

			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}

			for j := start; j <= i; j++ {
				result = append(result, runes[j])
			}

			if i+1 < len(runes) && runes[i+1] != ' ' && !goreloaded.Runponc(runes[i+1]) {
				result = append(result, ' ')
			}
			i++
		} else {
			result = append(result, r)
			i++
		}
	}

	return string(result)
}


func processTags(zrox []string) []string {
	for i := 0; i < len(zrox); i++ {
		switch zrox[i] {
		case "(cap)":
			if i != 0 {
				zrox[i-1] = goreloaded.Capitalize(zrox[i-1])
				zrox[i] = ""
				zrox = goreloaded.Cleanslice(zrox)
			}
		case "(up)":
			if i != 0 {
				zrox[i-1] = strings.ToUpper(zrox[i-1])
				zrox[i] = ""
				zrox = goreloaded.Cleanslice(zrox)
			}
		case "(low)":
			if i != 0 {
				zrox[i-1] = strings.ToLower(zrox[i-1])
				zrox[i] = ""
				zrox = goreloaded.Cleanslice(zrox)
			}
		case "(hex)":
			if i != 0 {
				num, err := strconv.ParseInt(zrox[i-1], 16, 64)
				if err != nil {
					fmt.Println("error converting hex:", err)
				} else {
					zrox[i-1] = strconv.Itoa(int(num))
					zrox[i] = ""
					zrox = goreloaded.Cleanslice(zrox)
					i--
				}
			}
		case "(bin)":
			if i != 0 {
				num, err := strconv.ParseInt(zrox[i-1], 2, 64)
				if err != nil {
					fmt.Println("error converting bin:", err)
				} else {
					zrox[i-1] = strconv.Itoa(int(num))
					zrox[i] = ""
					zrox = goreloaded.Cleanslice(zrox)
					i--
				}
			}
		case "(cap,":
			if i != 0 && i+1 < len(zrox) {
				end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
				for k := 1; k <= end; k++ {
					if i-k >= 0 {
						zrox[i-k] = goreloaded.Capitalize(zrox[i-k])
					}
				}
				zrox[i] = ""
				zrox[i+1] = ""
				zrox = goreloaded.Cleanslice(zrox)
			}
		case "(low,":
			if i != 0 && i+1 < len(zrox) {
				end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
				for k := 1; k <= end; k++ {
					if i-k >= 0 {
						zrox[i-k] = strings.ToLower(zrox[i-k])
					}
				}
				zrox[i] = ""
				zrox[i+1] = ""
				zrox = goreloaded.Cleanslice(zrox)
			}
		case "(up,":
			if i != 0 && i+1 < len(zrox) {
				end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
				for k := 1; k <= end; k++ {
					if i-k >= 0 {
						zrox[i-k] = strings.ToUpper(zrox[i-k])
					}
				}
				zrox[i] = ""
				zrox[i+1] = ""
				zrox = goreloaded.Cleanslice(zrox)
			}
		}

		
		if i == 0 {
			switch zrox[i] {
			case "(up)", "(cap)", "(low)", "(hex)", "(bin)":
				zrox[i] = ""
				zrox = goreloaded.Cleanslice(zrox)
			case "(up,", "(cap,", "(low,", "(hex,", "(bin,":
				if i+1 < len(zrox) {
					zrox[i] = ""
					zrox[i+1] = ""
					zrox = goreloaded.Cleanslice(zrox)
				}
			}
		}
	}
	return zrox
}


func writeOutput(filename string, zrox []string) error {
	var slice []byte
	for i, word := range zrox {
		slice = append(slice, []byte(word)...)
		if i != len(zrox)-1 {
			slice = append(slice, ' ')
		}
	}
	return os.WriteFile(filename, slice, 0o644)
}
