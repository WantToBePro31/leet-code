func findOriginalArray(changed []int) []int {
    var res []int
    sort.Slice(changed, func(i, j int) bool {
        return changed[i] < changed[j]
    })
    for i := 0; i < len(changed); i++ {
        if changed[i] == -1 {
            continue
        }
        ind := sort.Search(len(changed), func(j int) bool { return changed[j] >= changed[i]*2 && j != i })
        if ind < len(changed) && changed[ind] == changed[i]*2 && ind != i{
            res = append(res, changed[i])
            changed[i] = -1
            changed[ind] = -1
        } else {
            return []int{}
        }
        fmt.Println()
    }
    return res
}