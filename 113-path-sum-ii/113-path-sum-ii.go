/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchPath(root *TreeNode, sum, target int, path []int, nodePaths *[][]int) {
    if root == nil {
        return
    }
    sum += root.Val
    path = append(path, root.Val)
    if root.Left == nil && root.Right == nil && sum == target {
        (*nodePaths) = append((*nodePaths), append([]int{}, path...))
    }
    searchPath(root.Left, sum, target, path, nodePaths)
    searchPath(root.Right, sum, target, path, nodePaths)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    var nodePaths [][]int
    var path []int
    searchPath(root, 0, targetSum, path, &nodePaths)
    return nodePaths
}