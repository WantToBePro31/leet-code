/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    tmp := head
    count := 0
    for tmp != nil {
        count++
        tmp = tmp.Next
    }
    mid := count/2
    for mid > 0 {
        head = head.Next
        mid--
    }
    return head
}