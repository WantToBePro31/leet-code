func rotate(matrix [][]int) {
	l := len(matrix[0])
	tempMatrix := make([][]int, l)
	for i := range tempMatrix {
		tempMatrix[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			tempMatrix[j][l-1-i] = matrix[i][j]
		}
	}
    for i:= range tempMatrix {
        copy(matrix[i], tempMatrix[i])
    }
    fmt.Print("[")
    for i := 0; i < l; i++ {
        fmt.Print("[")
		for j := 0; j < l; j++ {
            if j != l-1 {
                fmt.Print(matrix[i][j], ",")
            } else {
                fmt.Print(matrix[i][j])
            }
		}
        if i != l-1 {
            fmt.Print("],")
        } else {
            fmt.Print("]]")
        }
	}
}