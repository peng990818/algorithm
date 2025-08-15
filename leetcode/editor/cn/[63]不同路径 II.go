//给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][
//n - 1]）。机器人每次只能向下或者向右移动一步。 
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。 
//
// 返回机器人能够到达右下角的不同路径数量。 
//
// 测试用例保证答案小于等于 2 * 10⁹。 
//
// 
//
// 示例 1： 
// 
// 
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
// 
//
// 示例 2： 
// 
// 
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// m == obstacleGrid.length 
// n == obstacleGrid[i].length 
// 1 <= m, n <= 100 
// obstacleGrid[i][j] 为 0 或 1 
// 
//
// Related Topics 数组 动态规划 矩阵 👍 1303 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//     if len(obstacleGrid) == 0 ||  len(obstacleGrid[0]) == 0 {
//         return 0
//     }
//     dp := make([][]int, len(obstacleGrid))
//     for i := range dp {
//         dp[i] = make([]int, len(obstacleGrid[0]))
//     }
//     if obstacleGrid[0][0] == 1 {
//         return 0
//     }
//     flagR, flagC := false, false
//     for i := 0;i<len(obstacleGrid);i++ {
//         for j := 0;j<len(obstacleGrid[0]);j++ {
//             if i == 0 {
//                 if flagR {
//                     dp[i][j] = 0
//                 } else {
//                     if obstacleGrid[i][j] == 1 {
//                         flagR = true
//                         dp[i][j] = 0
//                     } else {
//                         dp[i][j] = 1
//                     }
//                 }
//                 continue
//             }
//             if j == 0 {
//                 if flagC {
//                     dp[i][j] = 0
//                 } else {
//                     if obstacleGrid[i][j] == 1 {
//                         flagC = true
//                         dp[i][j] = 0
//                     } else {
//                         dp[i][j] = 1
//                     }
//                 }
//                 continue
//             }
//             if obstacleGrid[i][j] == 1 {
//                 dp[i][j] = 0
//             } else {
//                 dp[i][j] = dp[i-1][j] + dp[i][j-1]
//             }
//         }
//     }
//     return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
// }


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
dp := make([][]int, len(obstacleGrid))
for i := range dp {
dp[i] = make([]int, len(obstacleGrid[i]))
}
for i:=0;i<len(obstacleGrid[0]);i++ {
if obstacleGrid[0][i] == 1 {
break
}
dp[0][i] = 1
}
for i:=0;i<len(obstacleGrid);i++ {
if obstacleGrid[i][0] == 1 {
break
}
dp[i][0] = 1
}

for i:=1;i<len(obstacleGrid);i++ {
for j:=1;j<len(obstacleGrid[i]);j++ {
if obstacleGrid[i][j] == 1 {
continue
}
dp[i][j] = dp[i-1][j] + dp[i][j-1]
}
}
return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
//leetcode submit region end(Prohibit modification and deletion)
