//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。 
//
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = ["h","e","l","l","o"]
//输出：["o","l","l","e","h"]
// 
//
// 示例 2： 
//
// 
//输入：s = ["H","a","n","n","a","h"]
//输出：["h","a","n","n","a","H"] 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁵ 
// s[i] 都是 ASCII 码表中的可打印字符 
// 
//
// Related Topics 双指针 字符串 👍 903 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func reverseString(s []byte)  {
//     // 对称交换
//     m, n := len(s) / 2, len(s)-1
//     for i := 0; i < m; i++ {
//         s[i], s[n-i] = s[n-i], s[i]
//     }
//     return
// }

func reverseString(s []byte) {
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
s[i], s[j] = s[j], s[i]
}
}
//leetcode submit region end(Prohibit modification and deletion)
