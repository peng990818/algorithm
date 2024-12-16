//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。 
//
// 
//
// 注意： 
//
// 
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。 
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
// 
//
// 示例 2： 
//
// 
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
// 
//
// 示例 3: 
//
// 
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。 
//
// 
//
// 提示： 
//
// 
// m == s.length 
// n == t.length 
// 1 <= m, n <= 10⁵ 
// s 和 t 由英文字母组成 
// 
//
// 
//进阶：你能设计一个在 
//o(m+n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 3049 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 常规滑动窗口+哈希表
// func isCovered(cntS, cntT []int) bool {
//     for i:='A';i<='Z';i++ {
//         if cntS[i] < cntT[i] {
//             return false
//         }
//     }
//     for i:='a';i<='z';i++ {
//         if cntS[i] < cntT[i] {
//             return false
//         }
//     }
//     return true
// }
//
// func minWindow(s string, t string) string {
//     ansLeft, ansRight := -1, len(s)
//     var cntS, cntT [128]int
//     for _, c := range t {
//         cntT[c]++
//     }
//
//     left := 0
//     for right, c := range s {
//         cntS[c]++
//         for isCovered(cntS[:], cntT[:]) {
//             if right-left < ansRight-ansLeft {
//                 ansLeft, ansRight = left, right
//             }
//             cntS[s[left]]--
//             left++
//         }
//     }
//     if ansLeft < 0 {
//         return ""
//     }
//     return s[ansLeft:ansRight+1]
// }

// 优化
// func minWindow(s string, t string) string {
//     ansLeft, ansRight := -1, len(s)
//     cnt := [128]int{}
//     less := 0
//     for _, c := range t {
//         if cnt[c] == 0 {
//             less++ // 有less种字母的出现次数< t中的字母出现次数
//         }
//         cnt[c]++
//     }
//
//     left := 0
//     for right, c := range s {
//         cnt[c]--
//         if cnt[c] == 0 {
//             // 原来窗口内 c 的出现次数比 t 的少，现在一样多
//             less--
//         }
//         // 涵盖：所有字母的出现次数都是 >=
//         for less == 0 {
//             if right-left < ansRight-ansLeft {
//                 ansLeft, ansRight = left, right
//             }
//             x := s[left]
//             if cnt[x] == 0 {
//                 // x 移出窗口之前，检查出现次数，
//                 // 如果窗口内 x 的出现次数和 t 一样，
//                 // 那么 x 移出窗口后，窗口内 x 的出现次数比 t 的少
//                 less++
//             }
//             cnt[x]++
//             left++
//         }
//     }
//     if ansLeft < 0 {
//         return ""
//     }
//     return s[ansLeft:ansRight+1]
// }

// 常规版本
// func isCovered(cntS, cntT []int) bool {
//     for i:='A';i<='Z';i++ {
//         if cntS[i] < cntT[i] {
//             return false
//         }
//     }
//     for i:='a';i<='z';i++ {
//         if cntS[i] < cntT[i] {
//             return false
//         }
//     }
//     return true
// }
//
// func minWindow(s string, t string) string {
//     ansLeft, ansRight := -1, len(s)
//     var cntS, cntT [128]int
//     for _, c := range t {
//         cntT[c]++
//     }
//     left := 0
//     for right, c := range s {
//         cntS[c]++
//         for isCovered(cntS[:], cntT[:]) {
//             if right - left < ansRight - ansLeft {
//                 ansLeft, ansRight = left, right
//             }
//             cntS[s[left]]--
//             left++
//         }
//     }
//     if ansLeft < 0 {
//         return ""
//     }
//     return s[ansLeft:ansRight+1]
// }

func isCovered(cntT, cntS []int) bool {
    for i:='A';i<='Z';i++ {
        if cntS[i] < cntT[i] {
            return false
        }
    }
    for i:='a';i<='z';i++ {
        if cntS[i] < cntT[i] {
            return false
        }
    }
    return true
}

func minWindow(s string, t string) string {
    ansLeft, ansRight := -1, len(s)
    var cntT, cntS [128]int
    for _, c := range t {
        cntT[c]++
    }

    left := 0
    for right, c := range s {
        cntS[c]++
        for isCovered(cntT[:], cntS[:]) {
            if right-left < ansRight-ansLeft {
                ansLeft, ansRight = left, right
            }
            cntS[s[left]]--
            left++
        }
    }
    if ansLeft < 0 {
        return ""
    }
    return s[ansLeft:ansRight+1]
}

// func minWindow(s string, t string) string {
//     mpt := make(map[byte]int, len(t))
//     for i:=0;i<len(t);i++ {
//         mpt[t[i]]++
//     }
//     win := make(map[byte]int)
//     left := 0
//     resLeft, resRight := -1, len(s)
//     for right:=0;right < len(s);right++ {
//         win[s[right]]++
//         for isCovered(mpt, win) {
//             if right-left < resRight - resLeft {
//                 resLeft, resRight = left, right
//             }
//             win[s[left]]--
//             left++
//         }
//     }
//     if resLeft < 0 {
//         return ""
//     }
//     return s[resLeft: resRight+1]
// }
//
// // 是否包含
// func isCovered(mpt, win map[byte]int) bool {
//     for k, v := range mpt {
//         if val, ok := win[k]; ok {
//             if val < v {
//                 return false
//             }
//         } else {
//             return false
//         }
//     }
//     return true
// }
//leetcode submit region end(Prohibit modification and deletion)
