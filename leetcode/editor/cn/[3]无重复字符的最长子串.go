//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 
//
// 示例 2: 
//
// 
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 
//
// 示例 3: 
//
// 
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 10⁴ 
// s 由英文字母、数字、符号和空格组成 
// 
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10430 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func lengthOfLongestSubstring(s string) int {
//     if len(s) == 0 {
//         return 0
//     }
//     b := [128]byte{}
//     l, r, max := 0, 0, 1
//     for r < len(s) {
//         if b[s[r]] == 0 {
//             if max < r-l+1 {
//                 max = r-l+1
//             }
//         } else {
//             for l<=r {
//                 if s[l] == s[r] {
//                     b[s[l]] = 0
//                     l++
//                     break
//                 }
//                 b[s[l]] = 0
//                 l++
//             }
//         }
//         b[s[r]]++
//         r++
//     }
//     return max
// }

// func lengthOfLongestSubstring(s string) int {
//     // 哈希集合，记录每个字符是否出现过
//     m := map[byte]int{}
//     n := len(s)
//     rk, ans := -1, 0
//     for i:=0;i<n;i++ {
//         if i != 0 {
//             delete(m, s[i-1])
//         }
//         for rk+1<n && m[s[rk+1]] == 0 {
//             m[s[rk+1]]++
//             rk++
//         }
//         ans = max(ans, rk-i+1)
//     }
//     return ans
// }

// func lengthOfLongestSubstring(s string) int {
//     mp := make(map[byte]int)
//     left, right := 0, 0
//     res := 0
//     for right < len(s) {
//         mp[s[right]]++
//         for mp[s[right]] > 1 {
//             mp[s[left]]-=1
//             left++
//         }
//         res = max(res, right-left+1)
//         right++
//     }
//     return res
// }




func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    left, right := 0, 0
    res := 0
    for right < len(s) {
        mp[s[right]]++
        for mp[s[right]] > 1 {
            mp[s[left]]-=1
            left++
        }
        res = max(res, right-left+1)
        right++
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
