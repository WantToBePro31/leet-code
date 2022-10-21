func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func containsNearbyDuplicate(nums []int, k int) bool {
    mp1 := make(map[int]int)
    mp2 := make(map[int][]int)
    for ind, val := range nums {
        mp1[val]++
        if _, ok := mp2[val]; !ok {
            mp2[val] = append(mp2[val],ind)
        } else {
            if len(mp2[val]) != 2 {
                if abs(mp2[val][0], ind) <= k {
                    mp2[val] = append(mp2[val],ind)
                } else {
                    mp2[val][0] = ind
                }
            }
        }
    }
    for key, val := range mp1 {
        if val > 1 && len(mp2[key]) == 2 {
            return true
        }
    }
    return false
}