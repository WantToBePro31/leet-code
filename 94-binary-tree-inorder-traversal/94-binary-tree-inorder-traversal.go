/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func fillArray(root *TreeNode, nodes *[]int) {
    if root == nil {
        return
    }
    fillArray(root.Left, nodes)
    (*nodes) = append((*nodes), root.Val)
    fillArray(root.Right, nodes)
}

func inorderTraversal(root *TreeNode) []int {
    var nodes []int
    fillArray(root, &nodes)
    return nodes
}