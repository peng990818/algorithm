//给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。 
//
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 
//0（代表水）包围着。 
//
// 找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1
//,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0
//,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//输出: 6
//解释: 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。 
//
// 示例 2： 
//
// 
//输入: grid = [[0,0,0,0,0,0,0,0]]
//输出: 0 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 50 
// grid[i][j] is either 0 or 1 
// 
//
// 
//
// 注意：本题与主站 695 题相同： https://leetcode-cn.com/problems/max-area-of-island/ 
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 106 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func dfs(grid [][]int, i, j int,  visited [][]bool) int {
// res := 1
// if i-1 >= 0 && grid[i-1][j] == 1 && !visited[i-1][j] {
// visited[i-1][j] = true
// res += dfs(grid, i-1, j, visited)
// }
// if i+1 < len(grid) && grid[i+1][j] == 1 && !visited[i+1][j] {
// visited[i+1][j] = true
// res += dfs(grid, i+1, j, visited)
// }
// if j-1 >= 0 && grid[i][j-1] == 1 && !visited[i][j-1] {
// visited[i][j-1] = true
// res += dfs(grid, i, j-1, visited)
// }
// if j+1 < len(grid[0]) && grid[i][j+1] == 1 && !visited[i][j+1] {
// visited[i][j+1] = true
// res += dfs(grid, i, j+1, visited)
// }
// return res
// }
//
// func maxAreaOfIsland(grid [][]int) int {
// if len(grid) == 0  || len(grid[0]) == 0{
// return 0
// }
//
// res := 0
//
// visited := make([][]bool, len(grid))
// for i:=range visited {
// visited[i] = make([]bool, len(grid[0]))
// }
//
// for i:=0;i<len(grid);i++ {
// for j:=0;j<len(grid[0]);j++ {
// if grid[i][j] == 1 && !visited[i][j] {
// visited[i][j] = true
// res = max(res, dfs(grid, i, j, visited))
// }
// }
// }
// return res
// }

func bfs(grid [][]int, i, j int,  visited [][]bool) int {
res := 1
queue := make([][]int, 0)
queue = append(queue, []int{i, j})
for len(queue) > 0 {
r, c := queue[0][0], queue[0][1]
if r-1 >= 0 && grid[r-1][c] == 1 && !visited[r-1][c] {
visited[r-1][c] = true
res++
queue = append(queue, []int{r-1, c})
}
if r+1 < len(grid) && grid[r+1][c] == 1 && !visited[r+1][c] {
visited[r+1][c] = true
res++
queue = append(queue, []int{r+1, c})
}
if c-1 >= 0 && grid[r][c-1] == 1 && !visited[r][c-1] {
visited[r][c-1] = true
res++
queue = append(queue, []int{r, c-1})
}
if c+1 < len(grid[0]) && grid[r][c+1] == 1 && !visited[r][c+1] {
visited[r][c+1] = true
res++
queue = append(queue, []int{r, c+1})
}
queue = queue[1:]
}
return res
}

func maxAreaOfIsland(grid [][]int) int {
if len(grid) == 0  || len(grid[0]) == 0{
return 0
}
res := 0
visited := make([][]bool, len(grid))
for i:= range visited {
visited[i] = make([]bool, len(grid[0]))
}

for i:=0;i<len(grid);i++ {
for j:=0;j<len(grid[0]);j++ {
if grid[i][j] == 1 && !visited[i][j] {
visited[i][j] = true
res = max(res, bfs(grid, i, j, visited))
}
}
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
