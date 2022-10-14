func missingNumber(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    mp := make(map[int]int, nums[len(nums)-1])
    i := 0
    for _, val := range nums {
        mp[val] = 1
        if _, ok := mp[i]; !ok {
            return i
        }
        i++
    }
    return i
}