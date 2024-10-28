//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。 
//
// 说明：每次只能向下或者向右移动一步。 
//
// 
//
// 示例 1： 
// 
// 
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
// 
//
// 示例 2： 
//
// 
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
// 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 200 
// 0 <= grid[i][j] <= 200 
// 
//
// Related Topics 数组 动态规划 矩阵 👍 1737 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    for i:=0;i<len(grid);i++ {
        if i == 0 {
            for j := 1;j<len(grid[i]);j++ {
                grid[i][j] += grid[i][j-1]
            }
            continue
        }
        for j := 0;j<len(grid[i]);j++ {
            if j == 0 {
                grid[i][j] += grid[i-1][j]
                continue
            }
            if grid[i-1][j] < grid[i][j-1] {
                grid[i][j] += grid[i-1][j]
                continue
            }
            grid[i][j] += grid[i][j-1]
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}
//leetcode submit region end(Prohibit modification and deletion)
