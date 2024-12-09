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
// func spiralOrder(matrix [][]int) []int {
//     if len(matrix) == 0 || len(matrix[0]) == 0 {
//         return nil
//     }
//     res := make([]int, 0)
//     tR, tC, dR, dC := 0, 0, len(matrix)-1, len(matrix[0])-1
//     for tR <= dR && tC <= dC {
//         res = append(res, spiralEdge(matrix, tR, tC, dR, dC)...)
//         tR++
//         tC++
//         dR--
//         dC--
//     }
//     return res
// }
//
// func spiralEdge(matrix [][]int, tR, tC, dR, dC int) []int {
//     res := make([]int, 0)
//     if tR == dR {
//         // 仅有一行
//         for i:=tC;i<=dC;i++ {
//             res = append(res, matrix[tR][i])
//         }
//     } else if tC == dC {
//         // 仅有一列
//         for i:=tR;i<=dR;i++ {
//             res = append(res, matrix[i][tC])
//         }
//     } else {
//         i, j := tR, tC
//         for j < dC {
//             res = append(res, matrix[tR][j])
//             j++
//         }
//         for i < dR {
//             res = append(res, matrix[i][dC])
//             i++
//         }
//         for j > tC {
//             res = append(res, matrix[dR][j])
//             j--
//         }
//         for i > tR {
//             res = append(res, matrix[i][tC])
//             i--
//         }
//     }
//     return res
// }

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }
    tr, tc, dr, dc := 0, 0, len(matrix)-1, len(matrix[0])-1
    res := make([]int, 0, len(matrix)*len(matrix[0]))
    for tr <= dr && tc <= dc {
        res = append(res, spiralEdge(matrix, tr, tc, dr, dc)...)
        tr++
        tc++
        dr--
        dc--
    }
    return res
}

func spiralEdge(matrix [][]int, tr, tc, dr, dc int) []int {
    res := make([]int, 0)
    if tr == dr {
        for i:=tc;i<=dc;i++ {
            res = append(res, matrix[tr][i])
        }
    } else if tc == dc {
        for i:=tr;i<=dr;i++ {
            res = append(res, matrix[i][tc])
        }
    } else {
        i, j := tr, tc
        for j<dc {
            res = append(res, matrix[tr][j])
            j++
        }
        for i<dr {
            res = append(res, matrix[i][dc])
            i++
        }
        for j > tc {
            res = append(res, matrix[dr][j])
            j--
        }
        for i > tr {
            res = append(res, matrix[i][tc])
            i--
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
