//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 
//
// 示例 1： 
// 
// 
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 
//
// 示例 2： 
//
// 
//输入：l1 = [], l2 = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [], l2 = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 两个链表的节点数目范围是 [0, 50] 
// -100 <= Node.val <= 100 
// l1 和 l2 均按 非递减顺序 排列 
// 
//
// Related Topics 递归 链表 👍 3636 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     if list1 == nil {
//         return list2
//     }
//     if list2 == nil {
//         return list1
//     }
//     dummy := &ListNode{}
//     p3 := dummy
//     p1, p2 := list1, list2
//     for p1 != nil && p2 != nil {
//         if p1.Val < p2.Val {
//             p3.Next = p1
//             p1 = p1.Next
//         } else {
//             p3.Next = p2
//             p2 = p2.Next
//         }
//         p3 = p3.Next
//     }
//     if p1 != nil {
//         p3.Next = p1
//     }
//     if p2 != nil {
//         p3.Next = p2
//     }
//     return dummy.Next
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    dummy := &ListNode{}
    p := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            p.Next = list1
            list1 = list1.Next
        } else {
            p.Next = list2
            list2 = list2.Next
        }
        p = p.Next
    }
    if list1 != nil {
        p.Next = list1
    }
    if list2 != nil {
        p.Next = list2
    }
    return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
