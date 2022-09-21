/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchDepth(root *TreeNode, max *int, prev int) {
    if root == nil {
        return
    }
    if prev + 1 > *max {
        *max = prev + 1
    }
    searchDepth(root.Left, max, prev+1)
    searchDepth(root.Right, max, prev+1)
}

func fillNodes(root *TreeNode, nodesLevel *[][]int, level int) {
    if root == nil {
        return
    }
    (*nodesLevel)[level] = append((*nodesLevel)[level], root.Val)
    fillNodes(root.Left, nodesLevel, level+1)
    fillNodes(root.Right, nodesLevel, level+1)
}

func levelOrder(root *TreeNode) [][]int {
    max := 0
    searchDepth(root, &max, 0)
    nodesLevel := make([][]int, max)
    fillNodes(root, &nodesLevel, 0)
    return nodesLevel
}