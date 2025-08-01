//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。 
//
// 
//
// 示例 1： 
// 
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
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
// -1000 <= Node.val <= 1000 
// 
//
// Related Topics 树 广度优先搜索 二叉树 👍 2027 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func levelOrder(root *TreeNode) [][]int {
//     if root == nil {
//         return nil
//     }
//     res := make([][]int, 0)
//     queue := make([]*TreeNode, 0)
//     queue = append(queue, root)
//     for len(queue) > 0 {
//         tmp, newQueue := []int{}, make([]*TreeNode, 0)
//         for i:=0;i<len(queue);i++ {
//             tmp = append(tmp, queue[i].Val)
//             if queue[i].Left != nil {
//                 newQueue = append(newQueue, queue[i].Left)
//             }
//             if queue[i].Right != nil {
//                 newQueue = append(newQueue, queue[i].Right)
//             }
//         }
//         queue = newQueue
//         res = append(res, tmp)
//     }
//     return res
// }

// func levelOrder(root *TreeNode) [][]int {
//     if root == nil {
//         return nil
//     }
//     queue := make([]*TreeNode, 0)
//     res := make([][]int, 0)
//     queue = append(queue, root)
//     for len(queue) > 0 {
//         tmp := []int{}
//         newQueue := make([]*TreeNode, 0)
//         for i:=0;i<len(queue);i++ {
//             tmp = append(tmp, queue[i].Val)
//             if queue[i].Left != nil {
//                 newQueue = append(newQueue, queue[i].Left)
//             }
//             if queue[i].Right != nil {
//                 newQueue = append(newQueue, queue[i].Right)
//             }
//         }
//         queue = newQueue
//         res = append(res, tmp)
//     }
//     return res
// }

// func levelOrder(root *TreeNode) [][]int {
// if root == nil {
// return nil
// }
// queue := make([]*TreeNode, 0)
// queue = append(queue, root)
// res := make([][]int, 0)
// for len(queue) > 0 {
// var newQueue []*TreeNode
// var tmp []int
// for _, node := range queue {
// tmp = append(tmp, node.Val)
// if node.Left != nil {
// newQueue = append(newQueue, node.Left)
// }
// if node.Right != nil {
// newQueue = append(newQueue, node.Right)
// }
// }
// queue = newQueue
// res = append(res, tmp)
// }
// return res
// }

// 一个队列
func levelOrder(root *TreeNode) [][]int {
if root == nil {
return nil
}
queue := make([]*TreeNode, 0)
queue = append(queue, root)
res := make([][]int, 0)
for len(queue) > 0 {
size := len(queue)
var tmp []int
for i := 0; i < size; i++ {
tmp = append(tmp, queue[i].Val)
if queue[i].Left != nil {
queue = append(queue, queue[i].Left)
}
if queue[i].Right != nil {
queue = append(queue, queue[i].Right)
}
}
queue = queue[size:]
res = append(res, tmp)
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
