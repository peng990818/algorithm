//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
// 
//
// 示例 2： 
//
// 
//输入：root = [1]
//输出：[[1]]
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 2000] 内 
// -100 <= Node.val <= 100 
// 
//
// Related Topics 树 广度优先搜索 二叉树 👍 925 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    res := make([][]int, 0)
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    flag := true
    for len(queue) > 0 {
        var tmp []int
        newQueue := make([]*TreeNode, 0)
        for i:=0;i<len(queue);i++ {
            tmp = append(tmp, queue[i].Val)
            if queue[i].Left != nil {
                newQueue = append(newQueue, queue[i].Left)
            }
            if queue[i].Right != nil {
                newQueue = append(newQueue, queue[i].Right)
            }
        }
        flag = !flag
        queue = newQueue
        res = append(res, tmp)
        if flag {
            t, n := res[len(res)-1], len(res[len(res)-1])
            for i:=0;i<n/2;i++ {
                t[i], t[n-1-i] = t[n-1-i], t[i]
            }
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
