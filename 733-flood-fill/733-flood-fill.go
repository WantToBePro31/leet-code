func performFloodFill(image [][]int, sr, sc, color, row, col, pixel int) {
    if color == pixel || sr == -1 || sc == -1  || sr >= row || sc >= col {
        return
    }
    if image[sr][sc] == pixel {
        image[sr][sc] = color
        performFloodFill(image, sr-1, sc, color, row, col, pixel)
        performFloodFill(image, sr+1, sc, color, row, col, pixel)
        performFloodFill(image, sr, sc-1, color, row, col, pixel)
        performFloodFill(image, sr, sc+1, color, row, col, pixel)
    }
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    row, col := len(image), len(image[0])
    pixel := image[sr][sc]
    performFloodFill(image, sr, sc, color, row, col, pixel)
    return image
}