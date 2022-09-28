func getHint(secret string, guess string) string {
    bulls, cows := 0, 0
    var guess_tmp []rune
    count := make(map[rune]int)
    for ind, val := range secret {
        char := rune(guess[ind])
        if val != char {
            count[val]++
            guess_tmp = append(guess_tmp, char)
        } else {
            bulls++
        }
    }
    for _, val := range guess_tmp {
        if count[val] > 0 {
            cows++
            count[val]--
        }
    }
    s := fmt.Sprintf("%dA%dB", bulls, cows)
    return s
}