func numRollsToTarget(n int, k int, target int) int {
    if n * k < target {
         return 0
    }
    if n * k == target {
        return 1
    }
    modulo := 1000000007
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        var tmp []int
        for j := 0; j <= target; j++ {
            tmp = append(tmp, 0)
        }
        dp[i] = tmp
    }
    dp[0][0] = 1
    fmt.Println(dp)
    for i := 1; i <= n; i++ {
        for j := 1; j <= target; j++ {
            for a := 1; a <= k; a++ {
                if j >= a {
                    dp[i][j] = ((dp[i][j] % modulo) + (dp[i-1][j-a] % modulo)) % modulo
                }
            }
        }
    }
    return dp[n][target]
}