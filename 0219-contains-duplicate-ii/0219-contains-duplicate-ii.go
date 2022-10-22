func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func containsNearbyDuplicate(nums []int, k int) bool {
    mp1 := make(map[int]int)
    for ind, val := range nums {
        if _, ok := mp1[val]; !ok {
            mp1[val] = ind
        } else {
            if abs(mp1[val], ind) <= k {
                return true
            } else {
                mp1[val] = ind
            }
        }
    }
    return false
}