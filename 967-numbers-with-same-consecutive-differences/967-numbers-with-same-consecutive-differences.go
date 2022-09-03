func getNums(leading int, loop int, diff int, prev int, res *[]int) {
	if loop == 0 {
        for _, val := range *res {
			if val == leading {
				return
			}
    	}
		*res = append(*res, leading)
		return
	}

	if prev+diff > 9 && prev-diff < 0 {
		return
	}

	if prev+diff < 10 {
		getNums(leading+(prev+diff)*int(math.Pow(10, float64(loop-1))), loop-1, diff, prev+diff, res)
	}

	if prev-diff > -1 {
		getNums(leading+(prev-diff)*int(math.Pow(10, float64(loop-1))), loop-1, diff, prev-diff, res)
	}
}

func numsSameConsecDiff(n int, k int) []int {
	var res []int
	for i := 1; i < 10; i++ {
		leading := i * int(math.Pow(10, float64(n-1)))
		loop := n - 1
		getNums(leading, loop, k, i, &res)
	}
	return res
}