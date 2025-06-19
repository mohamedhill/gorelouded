package goreloaded

func StringToSlice(strclean string) []string {

    strclean += " "
    count := 0
    slice := []string{}

    for i := 0; i < len(strclean); i++ {
        if strclean[i] == ' ' {
            slice = append(slice, strclean[count:i])
            count = i+1
        }
    }
    return slice
    
}