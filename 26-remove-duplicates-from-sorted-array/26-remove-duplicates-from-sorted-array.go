func removeDuplicates(nums []int) int {
	remain := 0
    mp := make(map[int]int)
	for _, val := range nums {
		if _, found := mp[val]; found {
			mp[val]++
		} else {
			nums[remain] = val
			mp[val] = 1
			remain++
		}
	}
	return remain
}