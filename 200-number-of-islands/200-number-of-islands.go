func exploreIsland(grid [][]byte, sr, sc, row, col int) {
    if sr < 0 || sc < 0 || sr >= row || sc >= col {
        return
    }
    if grid[sr][sc] == '0' {
        return
    }
    grid[sr][sc] = '0'
    exploreIsland(grid, sr-1, sc, row, col)
    exploreIsland(grid, sr+1, sc, row, col)
    exploreIsland(grid, sr, sc-1, row, col)
    exploreIsland(grid, sr, sc+1, row, col)
}

func numIslands(grid [][]byte) int {
    row, col := len(grid), len(grid[0])
    islands := 0
    for sr := range grid {
        for sc, val := range grid[sr] {
            if val == '1' {
                exploreIsland(grid, sr, sc, row, col)
                islands++
            }
        }
    }
    return islands
}