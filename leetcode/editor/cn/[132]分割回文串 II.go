//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。 
//
// 返回符合要求的 最少分割次数 。 
//
// 
// 
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
// 
//
// 示例 2： 
//
// 
//输入：s = "a"
//输出：0
// 
//
// 示例 3： 
//
// 
//输入：s = "ab"
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 2000 
// s 仅由小写英文字母组成 
// 
//
// Related Topics 字符串 动态规划 👍 840 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
isPalindromic := make([][]bool, len(s))
for i := range isPalindromic {
isPalindromic[i] = make([]bool, len(s))
}

for i:=len(s)-1;i>=0;i-- {
for j:=i;j<len(s);j++ {
if s[i] == s[j] && (j-i<=1 || isPalindromic[i+1][j-1]) {
isPalindromic[i][j] = true
}
}
}

dp := make([]int, len(s))
for i:=range dp {
dp[i]= i
}
for i:=1;i<len(s);i++ {
if isPalindromic[0][i] {
dp[i] = 0
continue
}
for j:=0;j<i;j++ {
if isPalindromic[j+1][i] {
dp[i] = min(dp[i], dp[j]+1)
}
}
}
return dp[len(s)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
