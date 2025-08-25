//给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。 
//
// 测试用例保证结果在 32 位有符号整数范围内。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit 
//
// 示例 2： 
//
// 
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length, t.length <= 1000 
// s 和 t 由英文字母组成 
// 
//
// Related Topics 字符串 动态规划 👍 1366 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func numDistinct(s string, t string) int {
// dp[i][j] 表示以i-1为结尾的s可以随便删除元素，出现以j-1为结尾的字符串t的个数
dp := make([][]int, len(s)+1)
for i:=range dp {
dp[i] = make([]int, len(t)+1)
dp[i][0] = 1
}
for i:=1;i<=len(s);i++ {
for j:=1;j<=len(t);j++ {
if s[i-1] == t[j-1] {
dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
} else {
dp[i][j] = dp[i-1][j]
}
}
}
return dp[len(s)][len(t)]
}
//leetcode submit region end(Prohibit modification and deletion)
