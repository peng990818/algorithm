//给定两个字符串 s 和 t ，判断它们是否是同构的。 
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。 
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。 
//
// 
//
// 示例 1: 
//
// 
//输入：s = "egg", t = "add"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "foo", t = "bar"
//输出：false 
//
// 示例 3： 
//
// 
//输入：s = "paper", t = "title"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 
//
// 
// 1 <= s.length <= 5 * 10⁴ 
// t.length == s.length 
// s 和 t 由任意有效的 ASCII 字符组成 
// 
//
// Related Topics 哈希表 字符串 👍 738 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 我的方法：双哈希表
// func isIsomorphic(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }
//     // s2t的映射
//     mp := make(map[byte]byte, len(s))
//     // 元素是否被映射过
//     mp1 := make(map[byte]bool, len(s))
//     for i:=0;i<len(s);i++ {
//         v, ok := mp[s[i]]
//         if ok {
//             if v != t[i] {
//                 return false
//             }
//         } else {
//             if _, ok := mp1[t[i]]; ok {
//                 return false
//             }
//             mp[s[i]] = t[i]
//             mp1[t[i]] = true
//         }
//
//     }
//     return true
// }

// 官方
// func isIsomorphic(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }
//     s2t := make(map[byte]byte, len(s))
//     t2s := make(map[byte]byte, len(t))
//     for i:=0;i<len(s);i++ {
//         x, y := s[i], t[i]
//         if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
//             return false
//         }
//         s2t[x] = y
//         t2s[y] = x
//     }
//     return true
// }

func isIsomorphic(s string, t string) bool {
if len(s) != len(t) {
return false
}
mps := make(map[byte]byte, len(s))
mpt := make(map[byte]byte, len(t))
for i:=0;i<len(s);i++ {
v1, ok1 := mps[s[i]]
v2, ok2 := mpt[t[i]]
if ok1&&v1 != t[i] || ok2&&v2 != s[i] {
return false
}
mps[s[i]] = t[i]
mpt[t[i]] = s[i]
}
return true
}
//leetcode submit region end(Prohibit modification and deletion)
