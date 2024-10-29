//给你一个满足下述两条属性的 m x n 整数矩阵： 
//
// 
// 每行中的整数从左到右按非严格递增顺序排列。 
// 每行的第一个整数大于前一行的最后一个整数。 
// 
//
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
// 
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
// 
//
// 示例 2： 
// 
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 100 
// -10⁴ <= matrix[i][j], target <= 10⁴ 
// 
//
// Related Topics 数组 二分查找 矩阵 👍 977 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func searchMatrix(matrix [][]int, target int) bool {
//     if len(matrix) == 0 || len(matrix[0]) == 0 {
//         return false
//     }
//     m, n := len(matrix), len(matrix[0])
//     i := 0
//     for ;i<m;i++ {
//         if matrix[i][n-1] >= target {
//             break
//         }
//     }
//     if i == m {
//         return false
//     }
//     for j:=n-1;j>=0;j-- {
//         if matrix[i][j] == target {
//             return true
//         }
//     }
//     return false
// }

// 两次二分查找
// func searchMatrix(matrix [][]int, target int) bool {
//     // 按行二分，找到可能存在的行
//     row := sort.Search(len(matrix), func(i int) bool {
//         return matrix[i][0] > target
//     }) - 1
//     if row < 0 {
//         return false
//     }
//     col := sort.SearchInts(matrix[row], target)
//     return col < len(matrix[row]) && matrix[row][col] == target
// }

// 一次二分,把二维数组当成一维数组进行处理
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    i := sort.Search(m*n, func(i int) bool {
        return matrix[i/n][i%n] >= target
    })
    return i<m*n && matrix[i/n][i%n] == target
}
//leetcode submit region end(Prohibit modification and deletion)
