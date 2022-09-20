func maxProfit(prices []int) int {
    max := 0
    pivot := prices[0]
    for i := 1; i < len(prices); i++ {
        if pivot > prices[i] {
            pivot = prices[i] 
        } else if prices[i]-pivot > max {
            max = prices[i]-pivot
        }
    }
    return max
}