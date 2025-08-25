//给定两个单词 word1 和
// word2 ，返回使得
// word1 和 
// word2 相同所需的最小步数。 
//
// 每步 可以删除任意一个字符串中的一个字符。 
//
// 
//
// 示例 1： 
//
// 
//输入: word1 = "sea", word2 = "eat"
//输出: 2
//解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
// 
//
// 示例 2: 
//
// 
//输入：word1 = "leetcode", word2 = "etco"
//输出：4
// 
//
// 
//
// 提示： 
// 
//
// 
// 1 <= word1.length, word2.length <= 500 
// word1 和 word2 只包含小写英文字母 
// 
//
// Related Topics 字符串 动态规划 👍 740 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
dp := make([][]int, len(word1)+1)
for i:=range dp {
dp[i] = make([]int, len(word2)+1)
dp[i][0] = i
}
for j:=0;j<=len(word2);j++ {
dp[0][j] = j
}
for i:=1;i<=len(word1);i++ {
for j:=1;j<=len(word2);j++ {
if word1[i-1] == word2[j-1] {
dp[i][j] = dp[i-1][j-1]
} else {
dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
}
}
}
return dp[len(word1)][len(word2)]
}
//leetcode submit region end(Prohibit modification and deletion)
