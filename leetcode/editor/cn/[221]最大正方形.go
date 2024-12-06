//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。 
//
// 
//
// 示例 1： 
// 
// 
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
// 
//
// 示例 2： 
// 
// 
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
// 
//
// 示例 3： 
//
// 
//输入：matrix = [["0"]]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 300 
// matrix[i][j] 为 '0' 或 '1' 
// 
//
// Related Topics 数组 动态规划 矩阵 👍 1721 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func maximalSquare(matrix [][]byte) int {
//     dp := make([][]int, len(matrix))
//     maxSide := 0
//     for i:=0;i<len(matrix);i++ {
//         dp[i] = make([]int, len(matrix[i]))
//         for j:=0;j<len(matrix[i]);j++ {
//             dp[i][j] = int(matrix[i][j]-'0')
//             if dp[i][j] == 1 {
//                 maxSide = 1
//             }
//         }
//     }
//
//     for i:=1;i<len(matrix);i++ {
//         for j:=1;j<len(matrix[i]);j++ {
//             if dp[i][j] == 1 {
//                 dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
//                 if dp[i][j] > maxSide {
//                     maxSide = dp[i][j]
//                 }
//             }
//         }
//     }
//     return maxSide*maxSide
// }

func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    dp := make([][]int, len(matrix))
    maxSide := 0
    for i:=0;i<len(matrix);i++ {
        dp[i] = make([]int, len(matrix[0]))
        for j:=0;j<len(matrix[i]);j++ {
            dp[i][j] = int(matrix[i][j] - '0')
            if dp[i][j] == 1 {
                maxSide = 1
            }
        }
    }

    for i:=1;i<len(dp);i++ {
        for j:=1;j<len(dp[i]);j++ {
            if dp[i][j] == 1 {
                dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
                maxSide = max(maxSide, dp[i][j])
            }
        }
    }
    return maxSide*maxSide
}
//leetcode submit region end(Prohibit modification and deletion)
