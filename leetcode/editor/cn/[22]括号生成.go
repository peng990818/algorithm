//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 
//
// 示例 1： 
//
// 
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：["()"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 8 
// 
//
// Related Topics 字符串 动态规划 回溯 👍 3720 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
    res := make([]string, 0)
    var process func(path []byte, l, r int)
    process = func(path []byte, l, r int) {
        if len(path) == 2*n {
            res = append(res, string(path))
            return
        }
        if l < n {
            path = append(path, '(')
            process(path, l+1, r)
            path = path[:len(path)-1]
        }
        if r < l {
            path = append(path, ')')
            process(path, l, r+1)
            path = path[:len(path)-1]
        }
    }
    process(nil, 0, 0)
    return res
}

// func process(res *[]string, path []byte, lNum int, rNum int) {
//     if len(path) == 2*n {
//         tmp := string(path)
//         *res = append(*res, tmp)
//         return
//     }
//     if lNum > 0 {
//         path = append(path, '(')
//         lNum--
//         process(res, path, lNum, rNum)
//         path = path[:len(path)-1]
//     }
//     if lNum < rNum {
//         path = append(path, ')')
//         rNum--
//         process(res, path, lNum, rNum)
//         path = path[:len(path)-1]
//     }
// }
//leetcode submit region end(Prohibit modification and deletion)
