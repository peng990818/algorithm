//给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: head = [-10,-3,0,5,9]
//输出: [0,-3,9,-10,null,5]
//解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
// 
//
// 示例 2: 
//
// 
//输入: head = []
//输出: []
// 
//
// 
//
// 提示: 
//
// 
// head 中的节点数在[0, 2 * 10⁴] 范围内 
// -10⁵ <= Node.val <= 10⁵ 
// 
//
// Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 919 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMedian(left, right *ListNode) *ListNode {
    fast, slow := left, left
    for fast != right && fast.Next != right {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func sortedListToBST(head *ListNode) *TreeNode {
    return buildTree(head, nil)
}

func buildTree(left, right *ListNode) *TreeNode {
    if left == right {
        return nil
    }
    mid := getMedian(left, right)
    root := &TreeNode{Val: mid.Val}
    root.Left = buildTree(left, mid)
    root.Right = buildTree(mid.Next, right)
    return root
}
//leetcode submit region end(Prohibit modification and deletion)
