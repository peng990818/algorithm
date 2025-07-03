//给定一个非空的字符串
// s ，检查是否可以通过由它的一个子串重复多次构成。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abab"
//输出: true
//解释: 可由子串 "ab" 重复两次构成。
// 
//
// 示例 2: 
//
// 
//输入: s = "aba"
//输出: false
// 
//
// 示例 3: 
//
// 
//输入: s = "abcabcabcabc"
//输出: true
//解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
// 
//
// 
//
// 提示： 
//
// 
// 
//
// 
// 1 <= s.length <= 10⁴ 
// s 由小写英文字母组成 
// 
//
// Related Topics 字符串 字符串匹配 👍 1296 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
sub := []byte{}
for i := 0; i < len(s); i++ {
sub = append(sub, s[i])
if len(s)%len(sub) > 0 {
continue
}
if len(s) == len(sub) {
break
}
j := i + 1
bf := false
for j < len(s) {
flag := false
n := j
for m := 0; m < len(sub) && n < len(s); m, n = m+1, n+1 {
if sub[m] == s[n] {
continue
}
flag = true
break
}
if flag {
bf = true
break
}
j += len(sub)
}
if !bf {
return true
}
}
return false
}
//leetcode submit region end(Prohibit modification and deletion)
