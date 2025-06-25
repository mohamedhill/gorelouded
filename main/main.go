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
	data, err := os.ReadFile(filenames[1])
	
	clean := goreloaded.CleanStr(string(data))
	
	fmt.Println("str:",clean)
	runes := []rune(clean)
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

clean = string(result)
fmt.Println("str2:", clean)

zrox := goreloaded.StringToSlice(clean)
fmt.Println("str3:", zrox)

zrox = goreloaded.Cleanslice(zrox)


/* for l := 0; l < len(zrox); l++ {
	if l > 0 && strings.ContainsAny(zrox[l], ".,!?;:") {
		zrox[l-1] += zrox[l]
		zrox[l] = ""
	}
}


var final []string
for _, w := range zrox {
	if w != "" {
		final = append(final, w)
	}
} */
	

	
/* 
	var new []string
	var t string 
	 for k := 0; k < len(zrox); k++ {
		flag, runes:= goreloaded.Isponc(zrox[k])
		if flag {
			index := goreloaded.Index(zrox[k])
			if index == 0 && len(zrox[k]) == 1 {
				zrox[k]+= string(runes)+" " 
				
			
		
			 zrox[k] = zrox[k][1:]
				zrox[k-1] += string(runes)
				goreloaded.Cleanslice(zrox)
			} else if index == len(zrox[k])-1 && len(zrox[k]) > 1 && zrox[k][index-1] != '.' && zrox[k][index-1] != '?' && zrox[k][index-1] != ',' && zrox[k][index-1] != '!' &&
				zrox[k][index-1] != ';' && zrox[k][index-1] != ':' {
				fmt.Println(zrox)
			} else if index != 0 && index != len(zrox[k])-1 && len(zrox[k]) > 1 {
				new = append(new, zrox[k][:index])
				new = append(new, string(zrox[k][index]))
				new = append(new, zrox[k][index+1:])
				fmt.Println(new)
				t += zrox[k][:index]
				t += string(zrox[k][index])
				t += " "
				t += zrox[k][index+1:]
				new =goreloaded.StringToSlice(t)
				goreloaded.Cleanslice(new)
				r :=  new[0]+" "+goreloaded.Capitalize(new[1])
				fmt.Println(r)
				fmt.Println(new)

			}else if len(zrox[k]) == 1 {
				zrox[k] = ""
				zrox[k-1] += string(runes)
				fmt.Println("im here")

			}

		}
	}
  */

/* for l :=0 ;l < len(zrox);l++{


}
 */




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
		} else if i != 0 && zrox[i] == "(low," {
			end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
			for k := 1; k <= end; k++ {
				if i-k >= 0 {
					zrox[i-k] = strings.ToLower(zrox[i-k])
				}
			}
			zrox[i] = ""
			zrox[i+1] = ""
			zrox = goreloaded.Cleanslice(zrox)
		} else if i != 0 && zrox[i] == "(up," {
			end, _ := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
			for k := 1; k <= end; k++ {
				if i-k >= 0 {
					zrox[i-k] = strings.ToUpper(zrox[i-k])
				}
			}
			zrox[i] = ""
			zrox[i+1] = ""
			zrox = goreloaded.Cleanslice(zrox)
			// hna haydt hala dyal (up ...) fi index 0
		} else if i == 0 && zrox[i] == "(up)" || zrox[i] == "(cap)" || zrox[i] == "(low)" || zrox[i] == "(hex)" || zrox[i] == "(bin)" {
			zrox[i] = ""
			goreloaded.Cleanslice(zrox)
		} else if i == 0 && zrox[i] == "(up," || zrox[i] == "(cap," || zrox[i] == "(low," || zrox[i] == "(hex," || zrox[i] == "(bin," {
			zrox[i] = ""
			zrox[i+1] = ""
			goreloaded.Cleanslice(zrox)
		}
	}

	// covert slice string to byte and write the result file >>>>>>> string.join -or loop
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
