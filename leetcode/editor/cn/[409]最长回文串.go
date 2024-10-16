//给定一个包含大写字母和小写字母的字符串
// s ，返回 通过这些字母构造成的 最长的 回文串 的长度。 
//
// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。 
//
// 
//
// 示例 1: 
//
// 
//输入:s = "abccccdd"
//输出:7
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
// 
//
// 示例 2: 
//
// 
//输入:s = "a"
//输出:1
//解释：可以构造的最长回文串是"a"，它的长度是 1。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length <= 2000 
// s 只由小写 和/或 大写英文字母组成 
// 
//
// Related Topics 贪心 哈希表 字符串 👍 615 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
    // 统计频次
    mp := make(map[byte]int)
    for i := 0;i<len(s);i++ {
        if _, ok := mp[s[i]]; ok {
            mp[s[i]]++
            continue
        }
        mp[s[i]] = 1
    }
    if len(mp) == 1 {
        return len(s)
    }

    res, b := 0, false
    for _, v := range mp {
        // 偶数直接加
        if v % 2 == 0 {
            res+=v
            continue
        }
        b = true
        // 基数看情况
        res += v-1
    }
    if b {
        res++
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
