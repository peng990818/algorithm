//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。 
//
// 回文字符串 是正着读和倒过来读一样的字符串。 
//
// 子字符串 是字符串中的由连续字符组成的一个序列。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
// 
//
// 示例 2： 
//
// 
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa" 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 由小写英文字母组成 
// 
//
// Related Topics 双指针 字符串 动态规划 👍 1381 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
    n, res := len(s), 0
    for i:=0;i<2*n-1;i++ {
        l, r := i/2, i/2+i%2
        for l>=0&&r < n&&s[l] == s[r] {
            l--
            r++
            res++
        }
    }
    return res
    // todo 动态规划
}
//leetcode submit region end(Prohibit modification and deletion)
