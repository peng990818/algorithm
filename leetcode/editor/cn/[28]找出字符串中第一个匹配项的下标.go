//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
//如果 needle 不是 haystack 的一部分，则返回 -1 。 
//
// 
//
// 示例 1： 
//
// 
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。
// 
//
// 示例 2： 
//
// 
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= haystack.length, needle.length <= 10⁴ 
// haystack 和 needle 仅由小写英文字符组成 
// 
//
// Related Topics 双指针 字符串 字符串匹配 👍 2428 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func getNext(str string) []int {
// 1、初始化成-1
next := make([]int, 0, len(str))
for i := 0; i < len(str); i++ {
next = append(next, -1)
}
j := -1
for i := 1; i < len(str); i++ {
// 2、后缀不等于前缀
for j >= 0 && str[i] != str[j+1] {
j = next[j]
}
// 3、后缀等于前缀
if str[i] == str[j+1] {
j++
}
next[i] = j
}
return next
}

func strStr(haystack string, needle string) int {
if len(needle) == 0 {
return 0
}
next := getNext(needle)
j := -1
for i := 0; i < len(haystack); i++ {
for j >= 0 && haystack[i] != needle[j+1] {
j = next[j]
}
if haystack[i] == needle[j+1] {
j++
}
if j == len(needle)-1 {
return i - j
}
}
return -1
}
//leetcode submit region end(Prohibit modification and deletion)
