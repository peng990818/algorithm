//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。 
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。 
//
// 
// 
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
// 
// 
//
// 
//
// 示例 1： 
// 
// 
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：[["Q"]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 9 
// 
//
// Related Topics 数组 回溯 👍 2309 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
res := make([][]string, 0)
path := make([][]byte, n)
for i := 0; i < n; i++ {
path[i] = make([]byte, n)
}
for i := 0; i < n; i++ {
for j := 0; j < n; j++ {
path[i][j] = '.'
}
}
back(&res, path, 0, n)
return res
}

// 不能同行
// 不能同列
// 不能同斜线 （45度和135度角）
func isValid(row, col int, path [][]byte, n int) bool {
// 检查列
for i := 0; i < row; i++ {
if path[i][col] == 'Q' {
return false
}
}
// 同行无需检查，因为每一行只会有一个
// 45度
for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
if path[i][j] == 'Q' {
return false
}
}

// 135
for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
if path[i][j] == 'Q' {
return false
}
}
return true
}

func back(result *[][]string, path [][]byte, row int, n int) {
if row == n {
tmp := make([]string, n)
for i := 0; i < n; i++ {
tmp[i] = string(path[i])
}
*result = append(*result, tmp)
return
}
for col := 0; col < n; col++ {
if isValid(row, col, path, n) {
path[row][col] = 'Q'
back(result, path, row+1, n)
path[row][col] = '.'
}
}
}

//leetcode submit region end(Prohibit modification and deletion)
