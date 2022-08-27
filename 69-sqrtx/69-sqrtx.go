
func mySqrt(x int) int {
	n := 0
	for n*n < x{
		n++
	}
	if n*n > x {
		n--
	}
	return n
}