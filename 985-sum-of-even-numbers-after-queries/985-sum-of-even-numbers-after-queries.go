func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    sum := 0
    for _, val := range nums {
        if val % 2 == 0 {
            sum += val
        }
    }
    var res []int
    for i := 0; i < len(queries); i++ {
        val := queries[i][0]
        ind := queries[i][1]
        if nums[ind] % 2 == 0 {
            if (nums[ind] + val) % 2 == 0 {
                sum += val
            } else {
                sum -= nums[ind]
            }
            nums[ind] += val
        } else {
            if (nums[ind] + val) % 2 == 0 {
                sum += nums[ind] + val
            }
            nums[ind] += val
        }
        res = append(res, sum)
    }
    return res
}