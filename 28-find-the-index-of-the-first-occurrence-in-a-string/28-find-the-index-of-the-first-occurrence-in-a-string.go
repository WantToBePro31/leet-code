func strStr(haystack string, needle string) int {
	if needle == " " {
		return 0
	}
    var resInd int
	for i := range haystack {
		resInd = i
		ind := i
		gotIt := true
		for j := 0; j < len(needle); j++ {
			if ind == len(haystack) {
				gotIt = false
				break
			} else {
				if haystack[ind] == needle[j] {
					ind++
				} else {
					gotIt = false
					break
				}
			}
		}
		if gotIt {
			return resInd
		}
	}
	return -1
}