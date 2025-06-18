package gorelouded
func ToLower(s string) string {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'A' && sliceS[i] <= 'Z' {
			sliceS[i] += 32
		}
	}
	return string(sliceS)

}