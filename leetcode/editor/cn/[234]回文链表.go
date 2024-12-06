//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,2,2,1]
//输出：true
// 
//
// 示例 2： 
// 
// 
//输入：head = [1,2]
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点数目在范围[1, 10⁵] 内 
// 0 <= Node.val <= 9 
// 
//
// 
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
//
// Related Topics 栈 递归 链表 双指针 👍 1964 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverse(head *ListNode) *ListNode {
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
//
// func isPalindrome(head *ListNode) bool {
//     // 1、找到中点：快慢指针 slow为中点
//     fast, slow := head, head
//     for fast.Next != nil && fast.Next.Next != nil {
//         fast = fast.Next.Next
//         slow = slow.Next
//     }
//
//     // 2、反转后半部分的链表
//     rHead := reverse(slow.Next)
//
//     p1, p2 := head, rHead
//
//     // 3、前后遍历对比
//     res := true
//     for p1 != nil && p2 != nil {
//         if p1.Val != p2.Val {
//             res = false
//         }
//         p1 = p1.Next
//         p2 = p2.Next
//     }
//
//     // 翻转回去，保证原链表不变
//     slow.Next = reverse(rHead)
//     return res
// }

func reverse(head *ListNode) *ListNode {
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

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return false
    }
    // 1、中点
    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    rHead := reverse(slow.Next)
    p1, p2 := head, rHead
    res := true
    for p1 != nil && p2 != nil {
        if p1.Val != p2.Val {
            res = false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    slow.Next = reverse(rHead)
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
