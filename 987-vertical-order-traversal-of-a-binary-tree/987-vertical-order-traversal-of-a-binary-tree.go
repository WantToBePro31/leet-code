/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findDepth(root *TreeNode, mp2 *map[int]int, ind int, min *int, left int) {
    if root == nil {
        return
    }
    findDepth(root.Left, mp2, ind-1, min, left)
    if ind <= left {
        if (*mp2)[ind] == 0 {
            (*min)++
            (*mp2)[ind] = 1
        }
    }
    findDepth(root.Right, mp2, ind+1, min, left)
}

func defineArray(array *[][]int, root *TreeNode, mp *map[int]int, ind int) {
    if root == nil {
        return
    }
    defineArray(array, root.Left, mp, ind-1)
    if (*mp)[ind] == 0 {
        *array = append(*array, []int{})
        (*mp)[ind] = 1
    }
    defineArray(array, root.Right, mp, ind+1)
} 

func entry(array *[][]int, row *[][]int, col *[][]int, root *TreeNode, rows int, ind int) {
    if root == nil {
        return
    }
    if ind == -1 {
        return
    }
    (*array)[ind] = append((*array)[ind], root.Val)
    (*row)[ind] = append((*row)[ind], rows)
    (*col)[ind] = append((*col)[ind], ind)
    entry(array, row, col, root.Left, rows+1, ind-1)
    entry(array, row, col, root.Right, rows+1, ind+1)
}

func verticalTraversal(root *TreeNode) [][]int {
    vot := make([][]int, 0)
    mp1 := make(map[int]int)    
    mp2 := make(map[int]int)
    min := -1
    defineArray(&vot, root, &mp1, 1001)
    row := make([][]int, len(vot), cap(vot))
    col := make([][]int, len(vot), cap(vot))
    findDepth(root, &mp2, 1002, &min, 1002)
    entry(&vot, &row, &col, root, 1, min)
    for i := 0; i < len(vot); i++ {
        for j := 0; j < len(vot[i]); j++ {
            for k := j+1; k < len(vot[i]); k++ {
               if ((vot[i][j] > vot[i][k]) && (row[i][j] == row[i][k]) && (col[i][j] == col[i][k])) || (row[i][j] > row[i][k]) {
                   vot[i][j], vot[i][k] = vot[i][k], vot[i][j]                   
                   row[i][j], row[i][k] = row[i][k], row[i][j]
                   col[i][j], col[i][k] = col[i][k], col[i][j]
               } 
            }
        }
    }
    return vot
}