/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkBST(root *TreeNode, preLeft *TreeNode, preRight *TreeNode) bool {
    if root == nil {
        return true
    }
    if (preLeft != nil && preLeft.Val <= root.Val) || (preRight != nil && preRight.Val >= root.Val) {
        return false
    }
    return checkBST(root.Left, root, preRight) && checkBST(root.Right, preLeft, root)
}

func isValidBST(root *TreeNode) bool {
    return checkBST(root, nil, nil)
}