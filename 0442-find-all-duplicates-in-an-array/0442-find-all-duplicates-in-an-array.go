func findDuplicates(nums []int) []int {
    var res []int
    mp := make(map[int]int)
    for _, val := range nums {
        mp[val]++
        if mp[val] == 2 {
            res = append(res, val)
        } 
    }
    return res
}