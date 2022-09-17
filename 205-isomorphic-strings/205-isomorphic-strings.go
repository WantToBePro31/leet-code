func isIsomorphic(s string, t string) bool {
    fo := make(map[byte]byte)
    filled := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        if fo[s[i]] == byte(0) {
            if filled[t[i]] == 0 {
                fo[s[i]] = t[i]
                filled[t[i]] = 1
            } else {
                return false
            }
        } else if fo[s[i]] != t[i] {
            return false
        }
    }
    return true
}