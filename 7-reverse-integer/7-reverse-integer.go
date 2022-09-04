func reverse(x int) int {
    var minus bool
    if x < 0 {
        x *= -1
        minus = true
    }
    num := 0
    for x > 0 {
        num = num*10 + x%10
        x /= 10
    }
    if minus {
        num *= -1
    }
    limit := int(math.Pow(2,float64(31)))
    if num >= limit || num < (limit*-1) {
        return 0
    }
    return num
}