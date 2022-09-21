/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func countSize(head *ListNode) int {
    sz := 0
    for head != nil {
        sz++
        head = head.Next
    }
    return sz
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    sz := countSize(head)
    if sz == n {
        return head.Next
    }
    counter := sz - n - 1
    tmp := head
    for counter > 0 {
        tmp = tmp.Next
        counter--
    }
    tmp.Next = tmp.Next.Next
    return head
}