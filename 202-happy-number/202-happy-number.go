func countSum(n int) int {
    ans := 0
    for n != 0 {
        rem := n%10
        ans = ans + (rem*rem)
        n/=10
    }
    return ans
}

func isHappy(n int) bool {
    for n != 0 {
        n = countSum(n)
        if n == 1 || n == 7 {
            return true
        }
        if n/10 == 0 {
            return false
        }
    }
    return false
}