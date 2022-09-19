func twoSum(numbers []int, target int) []int {
    for ind, val := range numbers {
        tmp := target-val
        key := sort.Search(len(numbers), func(i int) bool {
            return numbers[i] >= tmp && i != ind
        })
        if key < len(numbers) && numbers[key] == tmp && ind != key {
            return []int{ind+1, key+1}
        }
    }
    return []int{-1,-1}
}