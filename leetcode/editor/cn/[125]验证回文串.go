//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。 
//
// 字母和数字都属于字母数字字符。 
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
// 
//
// 示例 2： 
//
// 
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
// 
//
// 示例 3： 
//
// 
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 2 * 10⁵ 
// s 仅由可打印的 ASCII 字符组成 
// 
//
// Related Topics 双指针 字符串 👍 770 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 我的方法
// 直接在原字符串改动
var delta byte = 'a' - 'A'
func isPalindrome(s string) bool {
    if len(s) <= 1 {
        return true
    }
    i, j := 0, len(s)-1
    for i < len(s) && j >= 0 {
        if !isOwn(s[i]) {
            i++
            continue
        }
        if !isOwn(s[j]) {
            j--
            continue
        }
        if !equal(s[i], s[j]) {
            return false
        }
        i++
        j--
    }
    return true
}

func isOwn(x byte) bool {
    return (x >= '0' && x <= '9') || (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')
}

func equal(a, b byte) bool {
    if a == b {
        return true
    }
    // a是数字或者b是数字
    if (a >= '0' && a <= '9') || (b >= '0' && b <= '9') {
        return false
    }
    // 大写字母和小写字母判等
    if a > b && a-b == delta  {
        return true
    }
    if b > a && b-a == delta  {
        return true
    }
    return false
}

// 官方 需要申请空间，使用底层库,时间上要更快一点
// func isPalindrome(s string) bool {
// var sgood string
// for i := 0; i < len(s); i++ {
// if isalnum(s[i]) {
// sgood += string(s[i])
// }
// }
//
// n := len(sgood)
// sgood = strings.ToLower(sgood)
// for i := 0; i < n/2; i++ {
// if sgood[i] != sgood[n-1-i] {
// return false
// }
// }
// return true
// }
//
// func isalnum(ch byte) bool {
// return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
// }

//leetcode submit region end(Prohibit modification and deletion)
