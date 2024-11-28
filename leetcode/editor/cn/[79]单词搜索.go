//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 
//
// 示例 1： 
// 
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = 
//"ABCCED"
//输出：true
// 
//
// 示例 2： 
// 
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = 
//"SEE"
//输出：true
// 
//
// 示例 3： 
// 
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = 
//"ABCB"
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == board.length 
// n = board[i].length 
// 1 <= m, n <= 6 
// 1 <= word.length <= 15 
// board 和 word 仅由大小写英文字母组成 
// 
//
// 
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？ 
//
// Related Topics 数组 字符串 回溯 矩阵 👍 1894 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// type pair struct{ x, y int }
// var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
// func exist(board [][]byte, word string) bool {
//     vis := make([][]bool, len(board))
//     for i := range vis {
//         vis[i] = make([]bool, len(board[0]))
//     }
//     for i:=0;i<len(board);i++ {
//         for j := 0;j<len(board[0]);j++ {
//             if process(board, word, i, j, 0, vis) {
//                 return true
//             }
//         }
//     }
//     return false
// }
//
// func process(board [][]byte, word string, i, j int, k int, vis [][]bool) bool {
//     if board[i][j] != word[k] {
//         // 剪枝：当前字符不匹配
//         return false
//     }
//     if k == len(word) - 1 {
//         return true
//     }
//         vis[i][j] = true
//         defer func() {
//             vis[i][j] = false
//         }()
//         for _, dir := range directions {
//             if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < len(board) && 0 <= newJ && newJ < len(board[0]) && !vis[newI][newJ] {
//                 if process(board, word, newI, newJ, k+1, vis) {
//                     return true
//                 }
//             }
//         }
//         return false
// }
type pair struct {
    x, y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    vis := make([][]bool, len(board))
    for i := range vis {
        vis[i] = make([]bool, len(board[0]))
    }
    for i:=0;i<len(board);i++ {
        for j:=0;j<len(board[0]);j++ {
            if process(board, word, i, j, 0, vis) {
                return true
            }
        }
    }
    return false
}

func process(board [][]byte, word string, i, j int, k int, vis [][]bool) bool {
    if word[k] != board[i][j] {
        return false
    }
    if k == len(word)-1 {
        return true
    }
    vis[i][j] = true
    defer func() {
        vis[i][j] = false
    }()
    for _, d := range directions {
        newI, newJ := i+d.x, j+d.y
        if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) && !vis[newI][newJ] {
            if process(board, word, newI, newJ, k+1, vis) {
                return true
            }
        }
    }
    return false
}
//leetcode submit region end(Prohibit modification and deletion)
