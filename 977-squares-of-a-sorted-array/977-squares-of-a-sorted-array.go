func sortedSquares(nums []int) []int {
    for ind, val := range nums {
        nums[ind] = val * val
    }
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    return nums
}