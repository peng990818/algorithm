//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。 
//
// 
//
// 示例 1： 
// 
// 
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
// 
//
// 示例 2： 
// 
// 
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 10 
// -100 <= matrix[i][j] <= 100 
// 
//
// Related Topics 数组 矩阵 模拟 👍 1777 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }
    res := make([]int, 0)
    tR, tC, dR, dC := 0, 0, len(matrix)-1, len(matrix[0])-1
    for tR <= dR && tC <= dC {
        res = append(res, spiralEdge(matrix, tR, tC, dR, dC)...)
        tR++
        tC++
        dR--
        dC--
    }
    return res
}

func spiralEdge(matrix [][]int, tR, tC, dR, dC int) []int {
    res := make([]int, 0)
    if tR == dR {
        // 仅有一行
        for i:=tC;i<=dC;i++ {
            res = append(res, matrix[tR][i])
        }
    } else if tC == dC {
        // 仅有一列
        for i:=tR;i<=dR;i++ {
            res = append(res, matrix[i][tC])
        }
    } else {
        i, j := tR, tC
        for j < dC {
            res = append(res, matrix[tR][j])
            j++
        }
        for i < dR {
            res = append(res, matrix[i][dC])
            i++
        }
        for j > tC {
            res = append(res, matrix[dR][j])
            j--
        }
        for i > tR {
            res = append(res, matrix[i][tC])
            i--
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
