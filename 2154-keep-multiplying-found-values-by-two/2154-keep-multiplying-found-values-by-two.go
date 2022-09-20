func findFinalValue(nums []int, original int) int {
    mp := make(map[int]int)
    for _, val := range nums {
        mp[val]++
    }
    for {
        if mp[original] > 0 {
            original *= 2
        } else {
            break
        }
    }
    return original
}