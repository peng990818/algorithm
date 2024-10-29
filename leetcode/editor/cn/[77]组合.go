//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//
// 示例 2： 
//
// 
//输入：n = 1, k = 1
//输出：[[1]]
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 20 
// 1 <= k <= n 
// 
//
// Related Topics 回溯 👍 1689 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
    if n == 0 || k == 0 {
        return nil
    }
    res := make([][]int, 0)
    path := make([]int, 0)
    process1(&res, path, 1, n, k)
    return res
}

func process(res *[][]int, path []int, l, r int, k int) {
    if len(path) + (r-l+1) < k {
        return
    }
    if len(path) == k {
        tmp := make([]int, k)
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    // 选择这个数
    path = append(path, l)
    l++
    process(res, path, l, r, k)
    path = path[:len(path)-1]
    process(res, path, l, r, k)
}

func process1(res *[][]int, path []int, l, r int, k int) {
    if len(path) == k {
        tmp := make([]int, k)
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    for l <= r {
        if len(path) + (r-l+1) < k {
            break
        }
        path = append(path, l)
        l++
        process(res, path, l, r, k)
        path = path[:len(path)-1]
    }
}
//leetcode submit region end(Prohibit modification and deletion)
