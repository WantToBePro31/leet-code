func singleNumber(nums []int) int {
    mp := make(map[int]int)
    for _, val := range nums {
        if mp[val] < 1 {
            mp[val]++
        } else {
            delete(mp, val)
        }
    }
    for val, _ := range mp {
        return val
    }
    
    return 0
}