/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recur(node *TreeNode, maxVal int, good *int) {
    if node != nil {
        if node.Val >= maxVal {
            maxVal = node.Val
            *good++
        }
    
        recur(node.Left, maxVal, good)
        recur(node.Right, maxVal, good)
    }
}

func goodNodes(root *TreeNode) int {
    init := 1
    good := &init
    recur(root.Left, root.Val, good)   
    recur(root.Right, root.Val, good)
    return *good
}