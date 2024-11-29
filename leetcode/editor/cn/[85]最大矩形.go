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
// func maximalRectangle(matrix [][]byte) (ans int) {
// if len(matrix) == 0 {
// return
// }
// m, n := len(matrix), len(matrix[0])
// left := make([][]int, m)
// for i, row := range matrix {
// left[i] = make([]int, n)
// for j, v := range row {
// if v == '0' {
// continue
// }
// if j == 0 {
// left[i][j] = 1
// } else {
// left[i][j] = left[i][j-1] + 1
// }
// }
// }
// for i, row := range matrix {
// for j, v := range row {
// if v == '0' {
// continue
// }
// width := left[i][j]
// area := width
// for k := i - 1; k >= 0; k-- {
// width = min(width, left[k][j])
// area = max(area, (i-k+1)*width)
// }
// ans = max(ans, area)
// }
// }
// return
// }
//
// func min(a, b int) int {
// if a < b {
// return a
// }
// return b
// }
//
// func max(a, b int) int {
// if a > b {
// return a
// }
// return b
// }

func maximalRectangle(matrix [][]byte) (ans int) {
    for i:=0;i<len(matrix);i++ {
        arr := make([]int, 0, len(matrix[0]))
        if i == 0 {
            for j:=0;j<len(matrix[0]);j++ {
                arr = append(arr, int(matrix[i][j]-'0'))
            }
            ans = max(ans, process(arr))
            continue
        }
        for j:=0;j<len(matrix[0]);j++ {
            if matrix[i][j] == '0' {
                arr = append(arr, 0)
            } else {
                tmp := int(matrix[i][j]-'0')
                for k:=i-1;k>=0;k-- {
                    if matrix[k][j] == '0' {
                        break
                    }
                    tmp += int(matrix[k][j]-'0')
                }
                arr = append(arr, tmp)
            }
            ans = max(ans, process(arr))
        }
    }
    return
}

func process(arr []int) int {
    n := len(arr)
    left, right := make([]int, n), make([]int, n)
    for i:=0;i<n;i++ {
        right[i] = n
    }

    stack := make([]int, 0)
    for i:=0;i<n;i++ {
        for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
            right[stack[len(stack)-1]] = i
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            left[i] = -1
        } else {
            left[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }
    res := 0
    for i:=0;i<n;i++ {
        res = max(res, (right[i]-left[i]-1)*arr[i])
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
