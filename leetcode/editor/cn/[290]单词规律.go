//给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。 
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。 
//
// 
//
// 示例1: 
//
// 
//输入: pattern = "abba", s = "dog cat cat dog"
//输出: true 
//
// 示例 2: 
//
// 
//输入:pattern = "abba", s = "dog cat cat fish"
//输出: false 
//
// 示例 3: 
//
// 
//输入: pattern = "aaaa", s = "dog cat cat dog"
//输出: false 
//
// 
//
// 提示: 
//
// 
// 1 <= pattern.length <= 300 
// pattern 只包含小写英文字母 
// 1 <= s.length <= 3000 
// s 只包含小写英文字母和 ' ' 
// s 不包含 任何前导或尾随对空格 
// s 中每个单词都被 单个空格 分隔 
// 
//
// Related Topics 哈希表 字符串 👍 675 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
    mp := make(map[byte]string)
    mpb := make(map[string]bool)
    j := 0
    for i := 0;i<len(pattern);i++ {
        if j >= len(s) {
            return false
        }
        tmp := make([]byte, 0)
        for j < len(s) {
            if s[j] == ' ' {
                j++
                break
            }
            tmp = append(tmp, s[j])
            j++
        }
        tWord := string(tmp)

        if word, ok := mp[pattern[i]]; ok {
            if word != tWord {
                return false
            }
        } else {
            if _, ok := mpb[tWord]; ok {
                return false
            }
            mp[pattern[i]] = tWord
            mpb[tWord] = true
        }
    }
    if j < len(s) {
        return false
    }
    return true
}
//leetcode submit region end(Prohibit modification and deletion)
