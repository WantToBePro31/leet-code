func minCost(colors string, neededTime []int) int {
    prev := ' '
    count := 0
    total := 0
    minTime := 0
    curTime := 0
    for ind, val := range colors {
        if val == prev {
            count++
            total += neededTime[ind]
            if neededTime[ind] > curTime {
                curTime = neededTime[ind]
            }
        } else {
            if count > 0 {
                minTime += total - curTime
            }
            total = neededTime[ind]
            curTime = neededTime[ind]
            count = 1
            prev = val
        }
    }
    if count > 0 {
        minTime += total - curTime
    }
    return minTime
}