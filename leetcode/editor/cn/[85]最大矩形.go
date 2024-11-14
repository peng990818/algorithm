//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。 
//
// 
//
// 示例 1： 
// 
// 
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
// 
//
// 示例 2： 
//
// 
//输入：matrix = [["0"]]
//输出：0
// 
//
// 示例 3： 
//
// 
//输入：matrix = [["1"]]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// rows == matrix.length 
// cols == matrix[0].length 
// 1 <= row, cols <= 200 
// matrix[i][j] 为 '0' 或 '1' 
// 
//
// Related Topics 栈 数组 动态规划 矩阵 单调栈 👍 1693 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) (ans int) {
if len(matrix) == 0 {
return
}
m, n := len(matrix), len(matrix[0])
left := make([][]int, m)
for i, row := range matrix {
left[i] = make([]int, n)
for j, v := range row {
if v == '0' {
continue
}
if j == 0 {
left[i][j] = 1
} else {
left[i][j] = left[i][j-1] + 1
}
}
}
for i, row := range matrix {
for j, v := range row {
if v == '0' {
continue
}
width := left[i][j]
area := width
for k := i - 1; k >= 0; k-- {
width = min(width, left[k][j])
area = max(area, (i-k+1)*width)
}
ans = max(ans, area)
}
}
return
}

func min(a, b int) int {
if a < b {
return a
}
return b
}

func max(a, b int) int {
if a > b {
return a
}
return b
}

//leetcode submit region end(Prohibit modification and deletion)
