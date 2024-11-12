//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。 
//
//
// 
//
// 
// 
// 示例 1： 
// 
// 
//
// 
//输入：n = 2
//输出：[0,1,1]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
// 
//
// 示例 2： 
//
// 
//输入：n = 5
//输出：[0,1,1,2,1,2]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
// 
//
// 
//
// 提示： 
//
// 
// 0 <= n <= 10⁵ 
// 
//
// 
//
// 进阶： 
//
// 
// 很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？ 
// 你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ） 
// 
//
// Related Topics 位运算 动态规划 👍 1348 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func countBits(n int) []int {
//     res := make([]int, n+1)
//     for i:=0;i<=n;i++ {
//         res[i] = onesCount(i)
//     }
//     return res
// }

// 令 x=x & (x−1)，该运算将 x 的二进制表示的最后一个 1 变成 0
// 因此，对 x 重复该操作，直到 x 变成 0，则操作次数即为 x 的「一比特数」。
// func onesCount(x int) (ones int) {
//     for ;x>0;x&=x-1 {
//         ones++
//     }
//     return
// }

func countBits(n int) []int {
    bits := make([]int, n+1)
    highBit := 0
    for i:=1;i<=n;i++ {
        if i&(i-1) == 0 {
            highBit = i
        }
        bits[i] = bits[i-highBit] + 1
    }
    return bits
}
//leetcode submit region end(Prohibit modification and deletion)
