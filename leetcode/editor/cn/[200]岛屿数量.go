//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。 
//
// 此外，你可以假设该网格的四条边均被水包围。 
//
// 
//
// 示例 1： 
//
// 
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 300 
// grid[i][j] 的值为 '0' 或 '1' 
// 
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2615 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 深度优先搜索
// func dfs(grid [][]byte, r, c int) {
//     nr, nc := len(grid), len(grid[0])
//     grid[r][c] = '0'
//     if (r-1 >= 0 && grid[r-1][c] == '1') {
//         dfs(grid, r-1, c)
//     }
//     if (r+1 < nr && grid[r+1][c] == '1') {
//         dfs(grid, r+1, c)
//     }
//     if (c-1 >= 0 && grid[r][c-1] == '1') {
//         dfs(grid, r, c-1)
//     }
//     if (c+1 < nc && grid[r][c+1] == '1') {
//         dfs(grid, r, c+1)
//     }
// }
//
// func numIslands(grid [][]byte) int {
//     if len(grid) == 0 {
//         return 0
//     }
//     res := 0
//     nr, nc := len(grid), len(grid[0])
//     for i:=0;i<nr;i++ {
//         for j:=0;j<nc;j++ {
//             if grid[i][j] == '1' {
//                 res++
//                 dfs(grid, i, j)
//             }
//         }
//     }
//     return res
// }

// func dfs(grid [][]byte, r, c int) {
//     nr, nc := len(grid), len(grid[0])
//     grid[r][c] = '0'
//     if r-1 >= 0 && grid[r-1][c] == '1' {
//         dfs(grid, r-1, c)
//     }
//     if r+1 < nr && grid[r+1][c] == '1' {
//         dfs(grid, r+1, c)
//     }
//     if c-1 >= 0 && grid[r][c-1] == '1' {
//         dfs(grid, r, c-1)
//     }
//     if c+1 < nc && grid[r][c+1] == '1' {
//         dfs(grid, r, c+1)
//     }
// }
//
// func numIslands(grid [][]byte) int {
//     if len(grid) == 0 || len(grid[0]) == 0 {
//         return 0
//     }
//     cnt := 0
//     for i:=0;i<len(grid);i++ {
//         for j:=0;j<len(grid[0]);j++ {
//             if grid[i][j] == '1' {
//                 dfs(grid, i, j)
//                 cnt++
//             }
//         }
//     }
//     return cnt
// }



// func dfs(grid [][]byte, i, j int) {
// nr, nc := len(grid), len(grid[0])
// grid[i][j] = '0'
// if i < nr-1 && grid[i+1][j] == '1' {
// dfs(grid, i+1, j)
// }
// if i>0 && grid[i-1][j] == '1' {
// dfs(grid,i-1,j)
// }
// if j < nc-1 && grid[i][j+1] == '1' {
// dfs(grid, i, j+1)
// }
// if j>0 && grid[i][j-1] == '1' {
// dfs(grid, i, j-1)
// }
// }
//
// func numIslands(grid [][]byte) int {
// if len(grid) == 0 || len(grid[0]) == 0 {
// return 0
// }
// var res int
// for i:=0;i<len(grid);i++ {
// for j:=0;j<len(grid[0]);j++ {
// if grid[i][j] == '1' {
// res++
// dfs(grid, i, j)
// }
// }
// }
// return res
// }

func bfs(grid [][]byte, i, j int) {
queue := make([][]int, 0)
queue = append(queue, []int{i, j})
grid[i][j] = '0'
for len(queue)>0{
r, c := queue[0][0], queue[0][1]
if r-1 >= 0 && grid[r-1][c] == '1' {
grid[r-1][c] = '0'
queue = append(queue, []int{r-1, c})
}
if r+1 < len(grid)&& grid[r+1][c] == '1' {
grid[r+1][c] = '0'
queue = append(queue, []int{r+1, c})
}
if c-1 >= 0 && grid[r][c-1] == '1' {
grid[r][c-1] = '0'
queue = append(queue, []int{r, c-1})
}
if c+1 < len(grid[0]) && grid[r][c+1] == '1' {
grid[r][c+1] = '0'
queue = append(queue, []int{r, c+1})
}
queue = queue[1:]
}
}

func numIslands(grid [][]byte) int {
if len(grid) == 0 || len(grid[0]) == 0 {
return 0
}
var res int
for i:=0;i<len(grid);i++ {
for j:=0;j<len(grid[0]);j++ {
if grid[i][j] == '1' {
res++
bfs(grid, i, j)
}
}
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
