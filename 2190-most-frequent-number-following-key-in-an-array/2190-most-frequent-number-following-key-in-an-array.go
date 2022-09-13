func mostFrequent(nums []int, key int) int {
    mp := make(map[int]int)
    max := 0
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] == key {
            mp[nums[i]]++
        }
        if mp[nums[i]] > count || (mp[nums[i]] == count && nums[i] > max) {
            max = nums[i]
            count = mp[nums[i]]
        }
    }
    return max
}