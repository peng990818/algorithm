//给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。 
//
// 美式键盘 中： 
//
// 
// 第一行由字符 "qwertyuiop" 组成。 
// 第二行由字符 "asdfghjkl" 组成。 
// 第三行由字符 "zxcvbnm" 组成。 
// 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：words = ["Hello","Alaska","Dad","Peace"]
//输出：["Alaska","Dad"]
// 
//
// 示例 2： 
//
// 
//输入：words = ["omk"]
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：words = ["adsdf","sfd"]
//输出：["adsdf","sfd"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= words.length <= 20 
// 1 <= words[i].length <= 100 
// words[i] 由英文字母（小写和大写字母）组成 
// 
//
// Related Topics 数组 哈希表 字符串 👍 260 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findWords(words []string) []string {
    t1 := "qwertyuiop"
    t2 := "asdfghjkl"
    t3 := "zxcvbnm"
    mp := make(map[rune]int, 26)
    for _, v := range t1 {
        mp[v] = 1
    }
    for _, v := range t2 {
        mp[v] = 2
    }
    for _, v := range t3 {
        mp[v] = 3
    }

    res := make([]string, 0)
    for _, word := range words {
        t, flag := 0, true
        word1 := strings.ToLower(word)
        for i, v := range word1 {
            if i != 0 && t != mp[v] {
                flag = false
                break
            } else {
                t = mp[v]
            }
        }
        if flag {
            res = append(res, word)
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
