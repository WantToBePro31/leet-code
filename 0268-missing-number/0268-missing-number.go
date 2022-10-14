func missingNumber(nums []int) int {
    max := 0
    tmp := 0
    for _, val := range nums {
        max++
        tmp += val
    }
    sum := max*(1+max)/2
    return sum-tmp
}