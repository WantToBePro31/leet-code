func twoSum(numbers []int, target int) []int {
    for ind, val := range numbers {
        for i := ind+1; i < len(numbers); i++ {
            if val + numbers[i] == target {
                return []int{ind+1, i+1}
            }
        }
    }
    return []int{-1,-1}
}