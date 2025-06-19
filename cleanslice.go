package goreloaded

func Cleanslice(s []string)[] string{
var clean []string
for i := 0 ; i <len(s); i++{
	if s[i]!=""{
clean = append(clean, s[i])
	}
}
return clean
}
