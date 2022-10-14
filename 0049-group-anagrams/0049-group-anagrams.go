func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }
    var res [][]string
    mp := make(map[string][]string)
    for _, val := range strs {
        str := []rune(val)
        sort.Slice(str, func(i int, j int) bool { 
            return str[i] < str[j]
        })
        mp[string(str)] = append(mp[string(str)], val)
    }
    
    for _, val := range mp {
        res = append(res, val)
    }
    return res
}