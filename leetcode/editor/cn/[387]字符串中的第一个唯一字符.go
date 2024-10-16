//给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。 
//
// 
//
// 示例 1： 
//
// 
//输入: s = "leetcode"
//输出: 0
// 
//
// 示例 2: 
//
// 
//输入: s = "loveleetcode"
//输出: 2
// 
//
// 示例 3: 
//
// 
//输入: s = "aabb"
//输出: -1
// 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length <= 10⁵ 
// s 只包含小写字母 
// 
//
// Related Topics 队列 哈希表 字符串 计数 👍 756 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 使用标准的map做hash表
// func firstUniqChar(s string) int {
//     mp := make(map[byte]int)
//     for i:=0;i<len(s);i++ {
//         if _, ok := mp[s[i]]; ok {
//             mp[s[i]]++
//             continue
//         }
//         mp[s[i]] = 1
//     }
//
//     for i:=0;i<len(s);i++ {
//         if mp[s[i]] == 1 {
//             return i
//         }
//     }
//     return -1
// }

// 使用数组代替hash表
func firstUniqChar(s string) int {
cnt := [26]int{}
for _, ch := range s {
cnt[ch-'a']++
}
for i, ch := range s {
if cnt[ch-'a'] == 1 {
return i
}
}
return -1
}

//leetcode submit region end(Prohibit modification and deletion)
