func wordPattern(pattern string, s string) bool {
    mp := make(map[rune]string)
    mpstr := make(map[string]int)
    strs := strings.Split(s, " ")
    if len(pattern) != len(strs) {
        return false
    }
    for ind, val := range pattern {
        if mp[val] == "" && mpstr[strs[ind]] == 0 {
            mp[val] = strs[ind]
            mpstr[strs[ind]] = 1
        } else {
            if strs[ind] != mp[val] {
                return false
            }
        }
    }
    return true
}