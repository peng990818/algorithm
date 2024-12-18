//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。 
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。 
//
// 问总共有多少条不同的路径？ 
//
// 
//
// 示例 1： 
// 
// 
//输入：m = 3, n = 7
//输出：28 
//
// 示例 2： 
//
// 
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
// 
//
// 示例 3： 
//
// 
//输入：m = 7, n = 3
//输出：28
// 
//
// 示例 4： 
//
// 
//输入：m = 3, n = 3
//输出：6 
//
// 
//
// 提示： 
//
// 
// 1 <= m, n <= 100 
// 题目数据保证答案小于等于 2 * 10⁹ 
// 
//
// Related Topics 数学 动态规划 组合数学 👍 2111 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 递归版本，数量大了必超时
// func uniquePaths(m int, n int) int {
//     i, j := 1, 1
//     res := 0
//     process(&res, i, j, m, n)
//     return res
// }
//
// func process(res *int, i, j, m, n int) {
//     if i == m && j == n {
//         (*res)++
//     }
//     if i < m && j < n {
//         i++
//         process(res, i, j, m, n)
//         i--
//         j++
//         process(res, i, j, m, n)
//     } else if i < m {
//         i++
//         process(res, i, j, m, n)
//     } else if j < n {
//         j++
//         process(res, i, j, m, n)
//     }
// }

// func uniquePaths(m int, n int) int {
//     if m == 0 || n == 0 {
//         return 0
//     }
//
//     dp := make([][]int, m)
//     for i := range dp {
//         if dp[i] == nil {
//             dp[i] = make([]int, n)
//             if i == 0 {
//                 for j := range dp[i] {
//                     dp[i][j] = 1
//                 }
//             }
//             dp[i][0] = 1
//         }
//     }
//
//     for i := 1;i<m;i++ {
//         for j := 1;j<n;j++ {
//             dp[i][j] = dp[i][j-1] + dp[i-1][j]
//         }
//     }
//     return dp[m-1][n-1]
// }

func uniquePaths(m int, n int) int {
    if m == 0 || n == 0 {
        return 0
    }
    dp := make([][]int, m)
    for i:=0;i<m;i++ {
        dp[i] = make([]int, n)
        for j := 0;j<n;j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 1
                continue
            }
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
