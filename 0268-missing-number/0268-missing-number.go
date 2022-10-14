func missingNumber(nums []int) int {
    max := 1
    for _, _ = range nums {
        max++
    }
    mp := make(map[int]int, max)
    for _, val := range nums {
        mp[val] = 1
    }
    i := 0
    for i < max {
        if _, ok := mp[i]; !ok {
            return i
        }
        i++
    }
    return max-1
}