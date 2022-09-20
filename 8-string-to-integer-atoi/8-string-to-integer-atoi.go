func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    regex, _ := regexp.Compile(`^[-|+]?[0-9]+`)

    num, _ := strconv.Atoi(regex.FindString(s))
    if num < math.MinInt32 {
        return math.MinInt32
    } else if num >  math.MaxInt32 {
        return math.MaxInt32
    } else {
        return num
    }
    return 0
}