package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"goreloaded"
)

func main() {
	// n9raw file hna >>> os READ FILE
	filenames := os.Args
	data, err := os.ReadFile(filenames[1])
	// nhaydo les espace >>> clean str
	clean := goreloaded.CleanStr(string(data))
	// nraj3oh slice
	zrox := goreloaded.StringToSlice(clean)
	// hna nbdaw n9ado lhalat li 3ndna >>>> ATOI - ITOA -TOUpper - Tolower -ParseINT
	for i := 0; i < len(zrox); i++ {
		if i != 0 && zrox[i] == "(cap)" {
			zrox[i-1] = goreloaded.Capitalize(zrox[i-1])
			zrox[i] = ""
			zrox = goreloaded.Cleanslice(zrox)
		} else if i != 0 && zrox[i] == "(up)" {
			zrox[i-1] = strings.ToUpper(zrox[i-1])
			zrox[i] = ""
			zrox = goreloaded.Cleanslice(zrox)

		} else if i != 0 && zrox[i] == "(low)" {
			zrox[i-1] = strings.ToLower(zrox[i-1])
			zrox[i] = ""
			zrox = goreloaded.Cleanslice(zrox)

		} else if i != 0 && zrox[i] == "(hex)" {
			num, err := strconv.ParseInt(zrox[i-1], 16, 64)
			if err != nil {
				fmt.Println("erorr converting", err)
			}
			zrox[i-1] = strconv.Itoa(int(num))
			zrox[i] = ""
			zrox = goreloaded.Cleanslice(zrox)
			i--

		} else if i != 0 && zrox[i] == "(bin)" {
			num, err := strconv.ParseInt(zrox[i-1], 2, 64)
			if err != nil {
				fmt.Println("erorr converting", err)
			}
			zrox[i-1] = strconv.Itoa(int(num))
			zrox[i] = ""
			zrox = goreloaded.Cleanslice(zrox)
			i--
       // hna n9ado blan dyal ar9am 
		} else if i != 0 && zrox[i] == "(cap," {
			end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
			for k := 1; k <= end; k++ {
				if i-k >= 0 {
					zrox[i-k] = goreloaded.Capitalize(zrox[i-k])
				}
			}
			zrox[i] = ""
			zrox[i+1] = ""
			zrox = goreloaded.Cleanslice(zrox)
		}else if i != 0 && zrox[i] == "(low," {
			end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
			for k := 1; k <= end; k++ {
				if i-k >= 0 {
					zrox[i-k] = strings.ToLower(zrox[i-k])
				}
			}
			zrox[i] = ""
			zrox[i+1] = ""
			zrox = goreloaded.Cleanslice(zrox)
		}else if i != 0 && zrox[i] == "(up," {
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
	zrox = goreloaded.Cleanslice(zrox)
	// covert slice string to byte and write thz result file >>>>>>> string.join -or loop
	/* combinedString := strings.Join(zrox, " ")
	   byteSlice := []byte(combinedString) */
	slice := []byte{}
	for i := 0; i < len(zrox); i++ {
		word := zrox[i]
		for j := 0; j < len(word); j++ {
			slice = append(slice, word[j])
		}
		if i != len(zrox)-1 {
			slice = append(slice, ' ')
		}
	}
	if err != nil {
		fmt.Println("error", err)
	} else {
		erre := os.WriteFile(filenames[2], []byte(slice), 0o644)
		if erre != nil {
			fmt.Println("error", err)
		}
	}
}
