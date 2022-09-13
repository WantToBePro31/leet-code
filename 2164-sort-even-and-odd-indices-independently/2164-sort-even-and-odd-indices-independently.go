func sortEvenOdd(nums []int) []int {
    if len(nums) == 1 || len(nums) == 2 {
        return nums
    }
    var (
        odd []int
        even []int
        final []int
    )
    for ind, val := range nums {
        if ind % 2 == 0 {
            even = append(even, val)
        } else {
            odd = append(odd, val)
        }
    }
    sort.Slice(odd, func(i, j int) bool {
        return odd[i] > odd[j]
    })
    sort.Slice(even, func(i, j int) bool {
        return even[i] < even[j]
    })
    for i := 0; i < len(nums)/2; i++ {      
        final = append(final, even[i])
        final = append(final, odd[i])  
    }
    if len(nums) % 2 == 1 {
        final = append(final, even[len(nums)/2])        
    }  
    return final
}