func longestPalindrome(s string) int {
    count, maxOdd, tmpOdd := 0, 0, 0
    firstIn := true
    mp := make(map[rune]int)
    for _, val := range s {
        mp[val]++
    }
    for _, n := range mp {
        if n % 2 == 0 {
            count += n
        } else if n % 2 == 1 {
            if n > maxOdd{
                if firstIn {
                    tmpOdd += n
                    firstIn = false
                } else {
                    tmpOdd += n-1
                }
                maxOdd = n
            } else {
                tmpOdd += n-1
            }
        }
    }
    return count + tmpOdd
}