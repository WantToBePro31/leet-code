func generate(numRows int) [][]int {
    pascal := make([][]int, numRows)
    if numRows > 0 {
        pascal[0] = append(pascal[0], 1)
    }
    if numRows > 1 {
        pascal[1] = append(pascal[1], 1)        
        pascal[1] = append(pascal[1], 1)
    }
    if numRows > 2 {
        for i := 2; i < numRows; i++ {
            pascal[i] = append(pascal[i], 1)
            for j := 0; j < len(pascal[i-1])-1; j++ {
                pascal[i] = append(pascal[i], pascal[i-1][j] + pascal[i-1][j+1])
            } 
            pascal[i] = append(pascal[i], 1)
        }
    }
    return pascal
}