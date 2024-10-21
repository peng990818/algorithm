//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。 
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。 
//
// 
//
// 示例 1: 
//
// 
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]] 
//
// 示例 2: 
//
// 
//输入: strs = [""]
//输出: [[""]]
// 
//
// 示例 3: 
//
// 
//输入: strs = ["a"]
//输出: [["a"]] 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 10⁴ 
// 0 <= strs[i].length <= 100 
// strs[i] 仅包含小写字母 
// 
//
// Related Topics 数组 哈希表 字符串 排序 👍 2021 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
    mp := make(map[string][]string)
    for _, str := range strs {
        // 将字符串转换为 []rune
        runes := []rune(str)

        // 对字符进行排序
        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })

        sortedStr := string(runes)
        if _, ok := mp[sortedStr] ;ok {
            mp[sortedStr] = append(mp[sortedStr], str)
            continue
        }
        mp[sortedStr] = []string{str}
    }
    res := make([][]string, 0, len(mp))
    for _, v := range mp {
        res = append(res, v)
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
