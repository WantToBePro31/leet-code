func reverseString(str string) string {
    tmpStr := ""
    for i := len(str)-1; i >= 0; i-- {
        tmpStr += string(str[i])
    }    
    return tmpStr
}

func reverseWords(s string) string {
    newS := strings.Split(s, " ")
    str := ""
    for i := 0; i < len(newS); i++ {
        str += reverseString(newS[i])
        if i != len(newS)-1 {
            str += " "
        }
    } 
    return str
}