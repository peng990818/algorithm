//给定一个字符串
// s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
// 
//
// 示例 2: 
//
// 
//输入： s = "Mr Ding"
//输出："rM gniD"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 5 * 10⁴ 
// 
// s 包含可打印的 ASCII 字符。 
// 
// s 不包含任何开头或结尾空格。 
// 
// s 里 至少 有一个词。 
// 
// s 中的所有单词都用一个空格隔开。 
// 
//
// Related Topics 双指针 字符串 👍 592 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) (res string) {
    if len(s) <= 1 {
        return s
    }
    var tmp []byte
    for i:=0;i<len(s);i++ {
        if s[i] != ' ' {
            tmp = append(tmp, s[i])
            continue
        }
        res += reverse(tmp) + " "
        tmp = nil
    }
    res += reverse(tmp)
    return
}

func reverse(bytes []byte) string {
    if len(bytes) == 0 {
        return ""
    }
    i, j := 0, len(bytes) - 1
    for i < j {
        bytes[i], bytes[j] = bytes[j], bytes[i]
        i++
        j--
    }
    return string(bytes)
}
//leetcode submit region end(Prohibit modification and deletion)
