//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。 
//
// 
// 如果剩余字符少于 k 个，则将剩余字符全部反转。 
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "abcdefg", k = 2
//输出："bacdfeg"
// 
//
// 示例 2： 
//
// 
//输入：s = "abcd", k = 2
//输出："bacd"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁴ 
// s 仅由小写英文组成 
// 1 <= k <= 10⁴ 
// 
//
// Related Topics 双指针 字符串 👍 682 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
s[i], s[j] = s[j], s[i]
}
}

// s = "abcdefg", k = 2  2k一组 "bacdfeg"
func reverseStr(s string, k int) string {
bs := []byte(s)
start := 0
for start < len(bs) {
end := start + k*2
if len(bs)-start >= k {
reverseString(bs[start : start+k])
} else {
reverseString(bs[start:])
}
start = end
}
return string(bs)
}
//leetcode submit region end(Prohibit modification and deletion)
