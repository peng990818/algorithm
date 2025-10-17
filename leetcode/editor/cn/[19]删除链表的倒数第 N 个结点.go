//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。 
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
// 
//
// 示例 2： 
//
// 
//输入：head = [1], n = 1
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：head = [1,2], n = 1
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 链表中结点的数目为 sz 
// 1 <= sz <= 30 
// 0 <= Node.val <= 100 
// 1 <= n <= sz 
// 
//
// 
//
// 进阶：你能尝试使用一趟扫描实现吗？ 
//
// Related Topics 链表 双指针 👍 2975 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     if head == nil {
//         return nil
//     }
//     p1, p2 := head, head
//     for n > 0 && p2 != nil {
//         p2 = p2.Next
//         n--
//     }
//     if p2 == nil {
//         return head.Next
//     }
//     for p2.Next != nil {
//         p1 = p1.Next
//         p2 = p2.Next
//     }
//     tmp := p1.Next.Next
//     p1.Next = tmp
//     return head
// }
//
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     if head == nil {
//         return head
//     }
//     p1, p2 := head, head
//     for n > 0 && p2 != nil {
//         p2 = p2.Next
//         n--
//     }
//     if p2 == nil {
//         // 删除第一个
//         return head.Next
//     }
//     for p2.Next != nil {
//         p1 = p1.Next
//         p2 = p2.Next
//     }
//     tmp := p1.Next.Next
//     p1.Next = tmp
//     return head
// }


func removeNthFromEnd(head *ListNode, n int) *ListNode {
dummy := &ListNode{}
dummy.Next = head
slow, fast := dummy, dummy
for n>0 && fast != nil {
fast = fast.Next
n--
}
fast = fast.Next
for fast != nil {
slow = slow.Next
fast = fast.Next
}
slow.Next = slow.Next.Next
return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
