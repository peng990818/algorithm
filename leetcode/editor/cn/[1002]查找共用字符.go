//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答
//案。
//
// 
//
// 示例 1： 
//
// 
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
// 
//
// 示例 2： 
//
// 
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= words.length <= 100 
// 1 <= words[i].length <= 100 
// words[i] 由小写英文字母组成 
// 
//
// Related Topics 数组 哈希表 字符串 👍 368 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func commonChars(words []string) []string {
    if len(words) == 0 {
        return nil
    }
    mp := make(map[rune][]int, 26)
    for i, w := range words {
        for _, c := range w {
            s, ok := mp[c]
            if ok {
                s[i]++
            } else {
                mp[c] = make([]int, len(words))
                mp[c][i] = 1
            }
        }
    }
    res := make([]string, 0)
    for k, v := range mp {
        for _, val := range v {
            if val == 0 {
                continue
            }
        }
        sort.Ints(v)
        for i:=0; i<v[0];i++ {
            res = append(res, string(k))
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
