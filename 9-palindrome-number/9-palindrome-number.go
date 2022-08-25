func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if(x >= 0 && x <= 9) {
        return true
    } else {
        tempX := x
        x2 := 0
        for tempX != 0 {
            add := tempX % 10
            x2 = x2 * 10 + add
            tempX /= 10
        }
        if x == x2 {
            return true
        }
    }
    return false
}