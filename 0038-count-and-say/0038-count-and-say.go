func recur(n int, mp map[string]string, count map[int]string) string {
    if n == 1 {
        return "1"
    }
    res := recur(n-1, mp, count)
	newNum := ""
    prev := rune(res[0])
	counter := 0
    for _, val := range res {
        if val != prev {
			str := fmt.Sprintf("%s %c", count[counter], prev)
			if _, ok := mp[str]; !ok {
				tmp := fmt.Sprintf("%d%c", counter, prev)
				mp[str] = tmp
			}
			newNum += mp[str]
			prev = val
			counter = 1
		} else {
			counter++
		}
    }
	str := fmt.Sprintf("%s %c", count[counter], prev)
	if _, ok := mp[str]; !ok {
		tmp := fmt.Sprintf("%d%c", counter, prev)
		mp[str] = tmp
	}
	newNum += mp[str]
	return newNum
}

func countAndSay(n int) string {
    mp := make(map[string]string)
	count := map[int]string {
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		10: "ten",
	}
    return recur(n, mp, count)
}