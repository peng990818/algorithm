//给你二叉树的根结点 root ，请你将它展开为一个单链表： 
//
// 
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。 
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。 
// 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [1,2,5,3,4,null,6]
//输出：[1,null,2,null,3,null,4,null,5,null,6]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 树中结点数在范围 [0, 2000] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？ 
//
// Related Topics 栈 树 深度优先搜索 链表 二叉树 👍 1748 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 1、前序遍历
// func flatten(root *TreeNode)  {
//     if root == nil {
//         return
//     }
//     list := preorder(root)
//     for i:=1;i<len(list);i++ {
//         pre, cur := list[i-1], list[i]
//         pre.Left, pre.Right = nil, cur
//     }
// }

func preorder(root *TreeNode) []*TreeNode {
    list := []*TreeNode{}
    if root != nil {
        list = append(list, root)
        list = append(list, preorder(root.Left)...)
        list = append(list, preorder(root.Right)...)
    }
    return list
}

// 2、前序遍历的同时，完成链表的连接
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    pre := (*TreeNode)(nil)
    for len(stack) > 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pre != nil {
            pre.Left, pre.Right = nil, cur
        }
        if cur.Right != nil {
            stack = append(stack, cur.Right)
        }
        if cur.Left != nil {
            stack = append(stack, cur.Left)
        }
        pre = cur
    }
}

//leetcode submit region end(Prohibit modification and deletion)
