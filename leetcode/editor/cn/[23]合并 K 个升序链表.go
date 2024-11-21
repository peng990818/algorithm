//给你一个链表数组，每个链表都已经按升序排列。 
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。 
//
// 
//
// 示例 1： 
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
// 
//
// 示例 2： 
//
// 输入：lists = []
//输出：[]
// 
//
// 示例 3： 
//
// 输入：lists = [[]]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// k == lists.length 
// 0 <= k <= 10^4 
// 0 <= lists[i].length <= 500 
// -10^4 <= lists[i][j] <= 10^4 
// lists[i] 按 升序 排列 
// lists[i].length 的总和不超过 10^4 
// 
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2913 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(list1, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    dummy := &ListNode{}
    p1, p2, p3 := list1, list2, dummy
    for p1 != nil && p2 != nil {
        if p1.Val < p2.Val {
            p3.Next = p1
            p1 = p1.Next
        } else {
            p3.Next = p2
            p2 = p2.Next
        }
        p3 = p3.Next
    }
    if p1 != nil {
        p3.Next = p1
    }
    if p2 != nil {
        p3.Next = p2
    }
    return dummy.Next
}

// 传统的两两合并
// func mergeKLists(lists []*ListNode) *ListNode {
//     if len(lists) == 0 {
//         return nil
//     }
//     tmp := lists[0]
//     for i:=1;i<len(lists);i++ {
//         tmp = merge(tmp, lists[i])
//     }
//     return tmp
// }

func mergeKLists(lists []*ListNode) *ListNode {
    return process(lists, 0, len(lists)-1)
}

func process(lists []*ListNode, l, r int) *ListNode {
    if l == r {
        return lists[l]
    }
    if l > r {
        return nil
    }
    mid := (l+r) >> 1
    return merge(process(lists, l, mid), process(lists, mid+1, r))
}
//leetcode submit region end(Prohibit modification and deletion)
