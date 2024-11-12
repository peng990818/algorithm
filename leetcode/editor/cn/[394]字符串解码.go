//给定一个经过编码的字符串，返回它解码后的字符串。 
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。 
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。 
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
// 
//
// 示例 2： 
//
// 
//输入：s = "3[a2[c]]"
//输出："accaccacc"
// 
//
// 示例 3： 
//
// 
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
// 
//
// 示例 4： 
//
// 
//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 30 
// 
// s 由小写英文字母、数字和方括号
// '[]' 组成 
// s 保证是一个 有效 的输入。 
// s 中所有整数的取值范围为
// [1, 300] 
// 
//
// Related Topics 栈 递归 字符串 👍 1870 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
stack := make([]byte, 0)
for i := 0; i < len(s); i++ {
if s[i] == ']' {
// 出栈
j, tmp := len(stack)-1, []byte{}
for stack[j] != '[' {
tmp = append(tmp, stack[j])
j--
}
j--
timesStr := []byte{}
for j >= 0 && stack[j] >= '0' && stack[j] <= '9' {
timesStr = append(timesStr, stack[j])
j--
}
reverse(timesStr)
times, _ := strconv.Atoi(string(timesStr))
stack = stack[:j+1]
reverse(tmp)
for t := 0; t < times; t++ {
stack = append(stack, tmp...)
}
} else {
// 入栈
stack = append(stack, s[i])
}
}
return string(stack)
}

func reverse(arr []byte) {
for i := 0; i < len(arr)/2; i++ {
arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
}
}
//leetcode submit region end(Prohibit modification and deletion)
