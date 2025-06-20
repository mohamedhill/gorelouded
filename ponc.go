package goreloaded

func Indexponc(s []string)int{
count := 0
for i := 0 ; i < len(s);i++{
	for j := 0 ; j< len(s[i]);j++{
		
		if s[i][j] == '.'{
			
			count = i
		}
		
	}
}
return count
}
