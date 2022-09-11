func divideArray(nums []int) bool {
    mp := make(map[int]int)
    for _, val := range nums {
        mp[val]++
    }
    for val, _ := range mp {
        if mp[val] % 2 == 1 {
            return false
        }
    }
    return true
}