func removeElement(nums []int, val int) int {
    remain := 0
    flag := 0
    for _, num := range nums {
        if num != val {
			nums[flag] = num
            remain++
            flag++
        }
    }
    return remain
}