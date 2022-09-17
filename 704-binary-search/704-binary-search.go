func search(nums []int, target int) int {
    ind := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= target
    })
    if ind < len(nums) && nums[ind] == target {
        return ind
    }
    return -1
}