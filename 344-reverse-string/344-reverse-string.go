func reverseString(s []byte)  {
    tmp := make([]byte, len(s))
    for ind, val := range s {
        tmp[len(s)-ind-1] = val
    }
    copy(s, tmp)
}