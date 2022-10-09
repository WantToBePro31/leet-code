/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isNodeExist(node, pivot *TreeNode, target int) bool {
    if node == nil {
        return false
    }
    if node != pivot {
        if node.Val == target {
            return true
        }
    }
    return isNodeExist(node.Left, pivot, target) || isNodeExist(node.Right, pivot, target)
}

func searchExistense(node, pivot *TreeNode, target int) bool {
    if node == nil {
        return false
    }
    newTarget := target-node.Val
    if isNodeExist(pivot, node, newTarget) {
        return true
    }
    return searchExistense(node.Left, pivot, target) || searchExistense(node.Right, pivot, target)
}

func findTarget(root *TreeNode, k int) bool {
    return searchExistense(root, root, k)
}