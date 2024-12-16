//给定一个已排序的链表的头
// head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。 
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,1,2]
//输出：[1,2]
// 
//
// 示例 2： 
// 
// 
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点数目在范围 [0, 300] 内 
// -100 <= Node.val <= 100 
// 题目数据保证链表已经按升序 排列 
// 
//
// Related Topics 链表 👍 1173 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func deleteDuplicates(head *ListNode) *ListNode {
//     if head == nil {
//         return nil
//     }
//     dummy := &ListNode{}
//     dummy.Next = head
//     p1, p2 := head, head.Next
//     for p2 != nil {
//         if p1.Val == p2.Val {
//             p2 = p2.Next
//             p1.Next = p2
//         } else {
//             p1 = p1.Next
//             p2 = p2.Next
//         }
//     }
//     return dummy.Next
// }

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    p1, p2 := head, head.Next
    for p2 != nil {
        if p1.Val == p2.Val {
            p2 = p2.Next
            p1.Next = p2
        } else {
            p1 = p1.Next
            p2 = p2.Next
        }
    }
    return head
}
//leetcode submit region end(Prohibit modification and deletion)
