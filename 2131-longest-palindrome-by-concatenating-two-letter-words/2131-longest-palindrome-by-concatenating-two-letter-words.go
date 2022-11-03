func longestPalindrome(words []string) int {
    mp := make(map[string]int)
    max := 0
    for _, val := range words {
        mp[val]++
        rev := fmt.Sprintf("%c%c", val[1], val[0])
        if val == rev && mp[val] == 2 {
            max += 4
            delete(mp, val)
        } else if val != rev {
            if _, ok := mp[rev]; ok && mp[rev] > 0 {
                max += 4
                mp[rev]--
                mp[val]--
            }
        }
    }
    for val, _ := range mp {
        if val[0] == val[1] {
            return max + 2
        }
    }
    return max
}