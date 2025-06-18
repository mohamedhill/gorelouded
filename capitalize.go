package gorelouded

func Capitalize(s string) string {
	sliceS := []rune(s)
	
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] <= 'Z' && sliceS[i] >= 'A' {
			sliceS[i] += 32
		}
	}
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] <= 'z' && sliceS[i] >= 'a' {
			sliceS[i] -= 32
		
	}
	
}
return string(sliceS)
}