//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。 
//
// 
// 
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
// 
//
// 示例 2： 
// 
// 
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
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
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内 
// -10⁵ <= Node.val <= 10⁵ 
// 
//
// 
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？ 
//
// Related Topics 链表 双指针 分治 排序 归并排序 👍 2397 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge(head1, head2 *ListNode) *ListNode {
    head := &ListNode{}
    p := head
    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            p.Next = head1
            head1 = head1.Next
        } else {
            p.Next = head2
            head2 = head2.Next
        }
        p = p.Next
    }
    if head1 != nil {
        p.Next = head1
    } else if head2 != nil {
        p.Next = head2
    }

    return head.Next
}

func sort(head, tail *ListNode) *ListNode {
    if head == nil {
        return head
    }
    if head.Next == tail {
        head.Next = nil
        return head
    }

    slow, fast := head, head
    for fast != tail {
        slow = slow.Next
        fast = fast.Next
        if fast != tail {
            fast = fast.Next
        }
    }
    mid := slow
    return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
   return sort(head, nil)
}
//leetcode submit region end(Prohibit modification and deletion)
