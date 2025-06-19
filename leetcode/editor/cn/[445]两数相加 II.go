//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。 
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。 
//
// 
//
// 示例1： 
//
// 
//
// 
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
// 
//
// 示例2： 
//
// 
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
// 
//
// 示例3： 
//
// 
//输入：l1 = [0], l2 = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 链表的长度范围为 [1, 100] 
// 0 <= node.val <= 9 
// 输入数据保证链表代表的数字无前导 0 
// 
//
// 
//
// 进阶：如果输入链表不能翻转该如何解决？ 
//
// Related Topics 栈 链表 数学 👍 755 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    // 压栈
    stack1, stack2 := []int{}, []int{}
    for l1 != nil {
        stack1 = append(stack1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        stack2 = append(stack2, l2.Val)
        l2 = l2.Next
    }
    i, j := len(stack1)-1, len(stack2)-1
    d := 0 // 进位
    var head *ListNode
    for i>=0 || j>=0 || d != 0 {
        num := d
        if i >= 0 {
            num += stack1[i]
            i--
        }
        if j >= 0 {
            num += stack2[j]
            j--
        }
        node := &ListNode{
            Val: num%10,
            Next: head,
        }
        head = node
        d = num/10
    }
    return head
}
//leetcode submit region end(Prohibit modification and deletion)
