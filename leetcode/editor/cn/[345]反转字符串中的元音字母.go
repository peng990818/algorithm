//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。 
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。 
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "IceCreAm" 
// 
//
// 输出："AceCreIm" 
//
// 解释： 
//
// s 中的元音是 ['I', 'e', 'e', 'A']。反转这些元音，s 变为 "AceCreIm". 
//
// 示例 2： 
//
// 
// 输入：s = "leetcode" 
// 
//
// 输出："leotcede" 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 3 * 10⁵ 
// s 由 可打印的 ASCII 字符组成 
// 
//
// Related Topics 双指针 字符串 👍 362 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 我的方法，双指针，内存占用较高
// var mp = map[byte]bool {
//     'a': true,
//     'e': true,
//     'i': true,
//     'o': true,
//     'u': true,
//     'A': true,
//     'E': true,
//     'I': true,
//     'O': true,
//     'U': true,
// }
// func reverseVowels(s string) string {
//     if len(s) <= 1 {
//         return s
//     }
//     bytes := []byte(s)
//     i, j := 0, len(bytes)-1
//     for i < j {
//         _, ok1 := mp[bytes[i]]
//         _, ok2 := mp[bytes[j]]
//         if ok1 && ok2 {
//             bytes[i], bytes[j] = bytes[j], bytes[i]
//             i++
//             j--
//         }
//         if !ok1 {
//             i++
//         }
//         if !ok2 {
//             j--
//         }
//     }
//     return string(bytes)
// }

// 官方
func reverseVowels(s string) string {
t := []byte(s)
n := len(t)
i, j := 0, n-1
for i < j {
for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
i++
}
for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
j--
}
if i < j {
t[i], t[j] = t[j], t[i]
i++
j--
}
}
return string(t)
}
//leetcode submit region end(Prohibit modification and deletion)
