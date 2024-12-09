//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。 
//
// 
//
// 示例 1： 
// 
// 
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 20 
// 
//
// Related Topics 数组 矩阵 模拟 👍 1345 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func generateMatrix(n int) [][]int {
//     if n == 0 {
//         return nil
//     }
//     res := make([][]int, n)
//     for i := range res {
//         res[i] = make([]int, n)
//     }
//
//     tR, tC, dR, dC := 0, 0, n-1, n-1
//     start := 1
//     for tR <= dR && tC <= dC {
//         start = process(res, tR, tC, dR, dC, start)
//         tR++
//         dR--
//         tC++
//         dC--
//     }
//     return res
// }
//
// func process(matrix [][]int, tR, tC, dR, dC int, start int) int {
//     if tR == dR {
//         for i:=tC;i<=dC;i++ {
//             matrix[tR][i] = start
//             start++
//         }
//     } else if tC == dC {
//         for i:=tR;i<=dR;i++ {
//             matrix[i][tC] = start
//             start++
//         }
//     } else {
//         i, j := tR, tC
//         for j < dC {
//             matrix[tR][j] = start
//             start++
//             j++
//         }
//         for i < dR {
//             matrix[i][dC] = start
//             start++
//             i++
//         }
//         for j > tC {
//             matrix[dR][j] = start
//             start++
//             j--
//         }
//         for i > tR {
//             matrix[i][tC] = start
//             start++
//             i--
//         }
//     }
//     return start
// }

func generateMatrix(n int) [][]int {
    if n <= 0 {
        return nil
    }
    matrix := make([][]int, n)
    for i:=0;i<n;i++ {
        matrix[i] = make([]int, n)
    }
    tr, tc, dr, dc := 0, 0, n-1, n-1
    start := 1
    for tr <= dr && tc <= dc {
        start = generateEdge(matrix, tr, tc, dr, dc, start)
        tr++
        tc++
        dr--
        dc--
    }
    return matrix
}

func generateEdge(matrix [][]int, tr, tc, dr, dc, start int) int {
    if tr == dr {
        for i:=tc;i<=dc;i++ {
            matrix[tr][i] = start
            start++
        }
    } else if tc == dc {
        for i:=tr;i<=dr;i++ {
            matrix[i][tc] = start
            start++
        }
    } else {
i, j := tr, tc
for j < dc {
matrix[tr][j] = start
start++
j++
}
for i < dr {
matrix[i][dc] = start
start++
i++
}
for j > tc {
matrix[dr][j] = start
start++
j--
}
for i > tr {
matrix[i][tc] = start
start++
i--
}
    }
    return start

}
//leetcode submit region end(Prohibit modification and deletion)
