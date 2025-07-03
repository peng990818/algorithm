//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
// 
// 
// 
// 
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
// 
//
// 示例 2： 
// 
// 
//输入：head = [1,2]
//输出：[2,1]
// 
//
// 示例 3： 
//
// 
//输入：head = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点的数目范围是 [0, 5000] 
// -5000 <= Node.val <= 5000 
// 
//
// 
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？ 
//
// Related Topics 递归 链表 👍 3692 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 三个指针循环遍历
// func reverseList(head *ListNode) *ListNode {
//     if head == nil {
//         return nil
//     }
//     p1, p2, p3 := (*ListNode)(nil), head, head.Next
//     for p3 != nil {
//         p2.Next = p1
//         p1 = p2
//         p2 = p3
//         p3 = p3.Next
//     }
//     p2.Next = p1
//     return p2
// }

// 递归写法
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     newHead := reverseList(head.Next)
//     head.Next.Next = head
//     head.Next = nil
//     return newHead
// }






// 双指针
// func reverseList(head *ListNode) *ListNode {
//     if head == nil {
//         return nil
//     }
//     p1, p2 := (*ListNode)(nil), head
//     for p2.Next != nil {
//         p3 := p2.Next
//         p2.Next = p1
//         p1 = p2
//         p2 = p3
//     }
//     p2.Next = p1
//     return p2
// }

// 递归实现
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next ==  nil {
//         return head
//     }
//     newHead := reverseList(head.Next)
//     head.Next.Next = head
//     head.Next = nil
//     return newHead
// }

// 双指针
// func reverseList(head *ListNode) *ListNode {
//     if head == nil {
//         return nil
//     }
//     p1, p2 := (*ListNode)(nil), head
//     for p2.Next != nil {
//         p3 := p2.Next
//         p2.Next = p1
//         p1 = p2
//         p2 = p3
//     }
//     p2.Next = p1
//     return p2
// }

// 递归
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     newHead := reverseList(head.Next)
//     head.Next.Next = head
//     head.Next = nil
//     return newHead
// }

func reverseList(head *ListNode) *ListNode {
if head == nil {
return nil
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
//leetcode submit region end(Prohibit modification and deletion)
