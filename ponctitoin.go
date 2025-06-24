package goreloaded
import "strings"

func Isponc(s string) (bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '?' || s[i] == '!' || s[i] == ';' || s[i] == ':' || s[i] == ',' {
			return true 
		}
	}
	return false 
}

func Runponc(s rune)bool{

if s== '.' || s == '?' || s == '!' || s == ';' || s == ':' || s == ',' {
			return true 

}
return false 
}


func Index(s string) int {
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '?' || s[i] == '!' || s[i] == ';' || s[i] == ':' || s[i] == ',' {
			index =strings.IndexRune(s, rune(s[i]))
		}
	}
	return index
}
 func Isflags(s string)bool{
	if s == "(up)"{
		return true
	}
	return false
 }