//给定一个单链表 L 的头节点 head ，单链表 L 表示为： 
//
// 
//L0 → L1 → … → Ln - 1 → Ln
// 
//
// 请将其重新排列后变为： 
//
// 
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → … 
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：head = [1,2,3,4]
//输出：[1,4,2,3] 
//
// 示例 2： 
//
// 
//
// 
//输入：head = [1,2,3,4,5]
//输出：[1,5,2,4,3] 
//
// 
//
// 提示： 
//
// 
// 链表的长度范围为 [1, 5 * 10⁴] 
// 1 <= node.val <= 1000 
// 
//
// Related Topics 栈 递归 链表 双指针 👍 1561 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reorderList(head *ListNode)  {
//     if head == nil {
//         return
//     }
//     arr := make([]*ListNode, 0)
//     for head != nil {
//         arr = append(arr, head)
//         head = head.Next
//     }
//     i, j := 0, len(arr)-1
//     for i<j {
//         arr[i].Next = arr[j]
//         i++
//         if i == j {
//             break
//         }
//         arr[j].Next = arr[i]
//         j--
//     }
//     arr[i].Next = nil
// }

func reverse(head *ListNode) *ListNode {
if head == nil || head.Next == nil {
return head
}
p1, p2 := (*ListNode)(nil), head
for p2.Next != nil {
p3 := p2.Next
p2.Next = p1
p1 = p2
p2 = p3
}
p2.Next = p1
return p2
}

func reorderList(head *ListNode)  {
if head == nil || head.Next == nil {
return
}
slow, fast := head, head
for fast.Next != nil && fast.Next.Next != nil {
slow = slow.Next
fast = fast.Next.Next
}
rHead := reverse(slow.Next)
slow.Next = nil
dummy := &ListNode{}
p := dummy
p1, p2 := head, rHead
flag := false
for p1 != nil || p2 != nil {
if !flag {
p.Next = p1
p1 = p1.Next
} else {
p.Next = p2
p2 = p2.Next
}
flag = !flag
p = p.Next
}
}
//leetcode submit region end(Prohibit modification and deletion)
