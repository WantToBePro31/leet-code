func countPrefixes(words []string, s string) int {
    ans := 0
    slen := len(s)
    for _, val := range words {
        if len(val) <= slen {
            flag := true
            sz := slen
            if len(val) < sz {
                sz = len(val)
            }
            for i := 0; i < sz; i++ {
                if val[i] != s[i] {
                    flag = false
                    break
                }
            }
            if flag {
                ans++
            }
        }
    }
    return ans
}