func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    if len(t) == 0 {
        return false
    }
    
    ind := 0
    for i := 0; i <len(t); i++ {
        if s[ind] == t[i] {
            ind++
        }
        if ind == len(s) {
            return true
        }
    }
    return false
}