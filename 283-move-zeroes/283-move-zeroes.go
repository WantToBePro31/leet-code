func partition(nums []int, low, high int) ([]int, int) {
	i := low
	for j := low; j < high; j++ {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
    }
    if nums[i] == 0 && nums[high] != 0 {
	    nums[i], nums[high] = nums[high], nums[i]
    }
	return nums, i
}

func quickSort(nums []int, low, high int) []int {
	if low < high {
		var pivot int
		nums, pivot = partition(nums, low, high)
		nums = quickSort(nums, low, pivot-1)
		nums = quickSort(nums, pivot+1, high)
    }
	return nums
}

func moveZeroes(nums []int)  {
    if len(nums) > 1 {
        copy(nums, quickSort(nums, 0, len(nums)-1))
    }
}