package goreloaded
import "strings"


func Removes(s []string)[]string{
for l := 0 ; l < len(s);l++{
for j := 0 ; j < len(s[l]) ;j++{
	if  l -1 > 0 && s[l][j]== '.' && s[l-1][j] == ' '{
	slc := strings.Fields(string(s))
	return strings.Join(slc ," ")
	
}

}

}
return s
}
