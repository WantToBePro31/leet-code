func isPalindrome(str string) bool {
    sz := len(str)
    if sz == 1 {
        return true
    }
    for i := 0; i < sz/2; i++ {
        if str[i] != str[sz-i-1] {
            return false
        }
    }
    return true
}

func longestPalindrome(s string) string {
    if len(s) == 1 {
        return s
    }
    ans := ""
    mp := make(map[string]int)
    for i := 0; i < len(s); i++ {
        tmps := ""
        buff := bytes.NewBufferString(tmps)
		buff.WriteByte(s[i])
		tmps = buff.String()
        for j := i+1; j < len(s); j++ {
            buff := bytes.NewBufferString(tmps)
			buff.WriteByte(s[j])
			tmps = buff.String()
            if mp[tmps] == 1 {
                continue
            }
            if isPalindrome(tmps) && len(tmps) > len(ans) {
                mp[tmps]++
                ans = tmps
            }
        }
    }
    if ans == "" {
        buff := bytes.NewBufferString(ans)
		buff.WriteByte(s[0])
		return buff.String()
    }
    return ans
}