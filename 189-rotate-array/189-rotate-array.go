func rotate(nums []int, k int)  {
    k %= len(nums)
    var tmp []int
    tmp = append(tmp, nums[len(nums)-k:]...)
    tmp = append(tmp, nums[:len(nums)-k]...)
    copy(nums, tmp) 
}