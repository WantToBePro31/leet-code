func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
        return len(strs[i]) < len(strs[j])
    })
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""
	pivot := strs[0]
	for i := 0; i < len(strs[0]); i++ {
		allSame := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != pivot[i] {
				allSame = false
				break
			}
		}
		if allSame {
			buff := bytes.NewBufferString(res)
			buff.WriteByte(pivot[i])
			res = buff.String()
		} else {
			return res
		}
	}
	return res
}