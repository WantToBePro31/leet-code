/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var list *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = list
        list = cur
        cur = next
    }
    return list
}