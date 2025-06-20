package goreloaded


func Isponc(s[]string)bool{
	for i := 0 ; i < len(s);i++{
		for j := 0 ; j <len(s[i]);j++{
      if s[i][j]== '.'{
		return  true
		
	  }
		}
	}
return false
}
