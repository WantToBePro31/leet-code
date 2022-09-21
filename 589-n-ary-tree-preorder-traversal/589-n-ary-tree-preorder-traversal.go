/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorderTraversal(root *Node, preorderNode *[]int) {
    if root == nil {
        return
    }
    *preorderNode = append(*preorderNode, root.Val)
    for i := 0; i < len(root.Children); i++ {
        preorderTraversal(root.Children[i], preorderNode)
    }
}

func preorder(root *Node) []int {
    var preorderNode []int
    preorderTraversal(root, &preorderNode)
    return preorderNode
}