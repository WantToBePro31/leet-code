func makeCombination(mp map[byte][]byte, digits string, ind, limit int, buffer bytes.Buffer, res *[]string) {
	if ind == limit {
		*res = append(*res, buffer.String())
		return
	}
	for _, val := range mp[digits[ind]] {
		tmp := buffer
		err := tmp.WriteByte(val)
		if err == nil {
			makeCombination(mp, digits, ind+1, limit, tmp, res)
		}
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	mp := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	var buffer bytes.Buffer
	var res []string
	makeCombination(mp, digits, 0, len(digits), buffer, &res)
	return res
}