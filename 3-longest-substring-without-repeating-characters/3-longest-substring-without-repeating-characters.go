func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 || len(s) == 1 {
        return len(s)
    }
    max, tmp, pivot, i := 0, 0, 0, 0
    mp := make(map[byte]int)
    for i < len(s) {
        if mp[s[i]] >= 1 {
            mp[s[pivot]]--
            if pivot != i {
                if tmp > max {
                    max = tmp
                }
                pivot++
                tmp -= 1
            } else {
                mp[s[i]]++
                i++
                tmp = 1
            }
        } else {
            mp[s[i]]++
            tmp++
            i++
        }
    }
    if tmp > max {
        max = tmp
    }
    return max
}