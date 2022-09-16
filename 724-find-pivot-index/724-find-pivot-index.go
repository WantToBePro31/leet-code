func pivotIndex(nums []int) int {
    sum := 0
    for _, val := range nums {
        sum += val
    }
    tmp := 0
    for ind, val := range nums {
        sum -= val
        if tmp == sum {
            return ind
        }
        tmp += val
    }
    return -1
}