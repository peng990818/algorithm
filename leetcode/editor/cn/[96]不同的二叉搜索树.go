//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。 
//
// 
//
// 示例 1： 
// 
// 
//输入：n = 3
//输出：5
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 19 
// 
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2565 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func numTrees(n int) int {
//     if n == 0 {
//         return 0
//     }
//     dp := make([]int, n+1)
//     dp[0], dp[1] = 1, 1
//     for i:=2;i<=n;i++ {
//         sum := 0
//         for j:=i-1;j>=0;j-- {
//             sum += dp[j]*dp[i-1-j]
//         }
//         dp[i] = sum
//     }
//     return dp[n]
// }

// func numTrees(n int) int {
//     if n == 0 {
//         return 0
//     }
//     dp := make([]int, n+1)
//     dp[0], dp[1] = 1, 1
//     for i:=2;i<n+1;i++ {
//         for j:=i-1;j>=0;j-- {
//             dp[i] += dp[j]*dp[i-1-j]
//         }
//     }
//     return dp[n]
// }

func numTrees(n int) int {
dp := make([]int, n+1)
dp[0] = 1
dp[1] = 1
for i := 2; i <= n; i++ {
for j := 0; j <= i-1; j++ {
dp[i] += dp[j] * dp[i-1-j]
}
}
return dp[n]
}
//leetcode submit region end(Prohibit modification and deletion)
