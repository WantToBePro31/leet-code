func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    score := 0
    low, high := 0, len(tokens)-1
    for low <= high {
        if power >= tokens[low] {
            score++
            power -= tokens[low]
            low++
        } else if score > 0 && high-low > 0{
            score--
            power += tokens[high]
            high--
        } else {
            return score
        }
    }
    return score
}