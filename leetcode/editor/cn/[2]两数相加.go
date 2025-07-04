//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。 
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。 
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。 
//
// 
//
// 示例 1： 
// 
// 
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
// 
//
// 示例 2： 
//
// 
//输入：l1 = [0], l2 = [0]
//输出：[0]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
// 
//
// 
//
// 提示： 
//
// 
// 每个链表中的节点数在范围 [1, 100] 内 
// 0 <= Node.val <= 9 
// 题目数据保证列表表示的数字不含前导零 
// 
//
// Related Topics 递归 链表 数学 👍 10906 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     if l1 == nil {
//         return l2
//     }
//     if l2 == nil {
//         return l1
//     }
//     dummy, flag := &ListNode{}, false
//     p := dummy
//     for l1 != nil && l2 != nil {
//         tmp := l1.Val + l2.Val
//         if flag {
//             tmp += 1
//             flag = false
//         }
//         if tmp >= 10 {
//             tmp -= 10
//             flag = true
//         }
//         p.Next = &ListNode{Val: tmp}
//         p = p.Next
//         l1 = l1.Next
//         l2 = l2.Next
//     }
//     for l1 != nil {
//         if flag {
//             l1.Val++
//             flag = false
//         }
//         if l1.Val >= 10 {
//             l1.Val -= 10
//             flag = true
//         }
//         p.Next = &ListNode{Val: l1.Val}
//         p = p.Next
//         l1 = l1.Next
//     }
//     for l2 != nil {
//         if flag {
//             l2.Val++
//             flag = false
//         }
//         if l2.Val >= 10 {
//             l2.Val -= 10
//             flag = true
//         }
//         p.Next = &ListNode{Val: l2.Val}
//         p = p.Next
//         l2 = l2.Next
//     }
//     if flag {
//         p.Next = &ListNode{Val: 1}
//     }
//     return dummy.Next
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{}
    d := 0
    p := dummy
    for l1 != nil || l2 != nil || d > 0 {
        num := d
        if l1 != nil {
            num += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            num += l2.Val
            l2 = l2.Next
        }
        p.Next = &ListNode{Val: num%10}
        p = p.Next
        d = num / 10
    }
    return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
