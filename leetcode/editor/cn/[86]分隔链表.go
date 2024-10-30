//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。 
//
// 你应当 保留 两个分区中每个节点的初始相对位置。 
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
// 
//
// 示例 2： 
//
// 
//输入：head = [2,1], x = 2
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点的数目在范围 [0, 200] 内 
// -100 <= Node.val <= 100 
// -200 <= x <= 200 
// 
//
// Related Topics 链表 双指针 👍 866 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 我的方法，较为麻烦，复杂度较高
// func partition(head *ListNode, x int) *ListNode {
//     if head == nil {
//         return nil
//     }
//     dummy := &ListNode{-201, head}
//     cur := dummy
//     for cur.Next != nil {
//         if cur.Next.Val < x {
//             cur = cur.Next
//             continue
//         }
//         p1 := cur.Next
//         for p1.Next != nil && p1.Next.Val >= x {
//             p1 = p1.Next
//         }
//         tmp := p1.Next
//         if p1.Next != nil {
//             p1.Next = p1.Next.Next
//             tmp.Next = cur.Next
//             cur.Next = tmp
//         }
//         cur = cur.Next
//     }
//     return dummy.Next
// }

// 官方解法
func partition(head *ListNode, x int) *ListNode {
    small := &ListNode{}
    smallHead := small
    large := &ListNode{}
    largeHead := large
    for head != nil {
        if head.Val < x {
            small.Next = head
            small = small.Next
        } else {
            large.Next = head
            large = large.Next
        }
        head = head.Next
    }
    large.Next = nil
    small.Next = largeHead.Next
    return smallHead.Next
}
//leetcode submit region end(Prohibit modification and deletion)
