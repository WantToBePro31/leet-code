func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}

func jump(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    dp := make([]int, len(nums))
    for ind := range dp {
        dp[ind] = 10001
    }
    dp[0] = 0
    for ind, val := range nums {
        for i := ind; i <= ind + val && i < len(nums); i++ {
            dp[i] = min(dp[i], dp[ind] + 1)
        }
    }
    return dp[len(nums)-1]
}