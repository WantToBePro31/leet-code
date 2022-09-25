func exploreIsland(grid [][]int, sr, sc, row, col int, area *int) {
    if sr < 0 || sc < 0 || sr >= row || sc >= col {
        return
    }
    if grid[sr][sc] == 0 {
        return
    }
    grid[sr][sc] = 0
    *area++
    exploreIsland(grid, sr-1, sc, row, col, area)
    exploreIsland(grid, sr+1, sc, row, col, area)
    exploreIsland(grid, sr, sc-1, row, col, area)
    exploreIsland(grid, sr, sc+1, row, col, area)
}

func maxAreaOfIsland(grid [][]int) int {
    row, col := len(grid), len(grid[0])
    max := 0
    for sr := range grid {
        for sc, val := range grid[sr] {
            if val == 1 {
                area := 0
                exploreIsland(grid, sr, sc, row, col, &area)
                if area > max {
                    max = area
                }
            }
        }
    }
    return max
}