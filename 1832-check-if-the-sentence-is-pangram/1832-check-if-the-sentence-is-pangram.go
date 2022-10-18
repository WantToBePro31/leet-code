func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }
    mp := make(map[rune]int)
    count := 0
    for _, char := range sentence {
        if _, ok := mp[char]; !ok {
            count++
            mp[char]++
        } 
    }
    if count >= 26 {
        return true
    }
    return false
}