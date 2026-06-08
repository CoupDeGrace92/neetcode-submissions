/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    order := make(map[int]*ListNode)
    if head == nil || head.Next == nil {
        return
    }
    current := head
    i := 0
    for current != nil {
        order[i] = current
        current = current.Next
        i++
    }

    front := false
    k := len(order) - 1
    current = order[0]
    for j := 1; j <= k; {
        if front {
            current.Next = order[j]
            current = order[j]
            j++
        } else {
            current.Next = order[k]
            current = order[k]
            k--
        }
        front = !front
    }
    current.Next = nil
}

//First thought - hash map with positions - then we can order them, so two pass
