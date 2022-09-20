func reverseString(s []byte)  {
    var tmp []byte
    for i := len(s)-1; i >= 0; i-- {
        tmp = append(tmp, s[i])
    }
    copy(s, tmp)
}