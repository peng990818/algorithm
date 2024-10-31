//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
// 
//
// 示例 1： 
// 
// 
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
// 
//
// 示例 2： 
//
// 
//输入：head = [5], left = 1, right = 1
//输出：[5]
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点数目为 n 
// 1 <= n <= 500 
// -500 <= Node.val <= 500 
// 1 <= left <= right <= n 
// 
//
// 
//
// 进阶： 你可以使用一趟扫描完成反转吗？ 
//
// Related Topics 链表 👍 1867 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 反转链表
func reverse(head *ListNode) {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right {
        return head
    }
    dummy := &ListNode{}
    dummy.Next = head
    // 1、找到左边前一个节点
    pre := dummy
    for i := 0; i < left-1;i++ {
        pre = pre.Next
    }

    // 2、找到结束节点
    rightNode := pre
    for i:=0;i<right-left+1;i++ {
        rightNode = rightNode.Next
    }

    // 3、制作新的链表
    leftNode := pre.Next
    curr := rightNode.Next

    pre.Next = nil
    rightNode.Next = nil

    reverse(leftNode)

    pre.Next = rightNode
    leftNode.Next = curr
    return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
