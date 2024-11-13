//给你一个非负整数数组 nums 和一个整数 target 。 
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ： 
//
// 
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。 
// 
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], target = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 20 
// 0 <= nums[i] <= 1000 
// 0 <= sum(nums[i]) <= 1000 
// -1000 <= target <= 1000 
// 
//
// Related Topics 数组 动态规划 回溯 👍 2036 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 回溯
// func findTargetSumWays(nums []int, target int) int {
//     var res int
//     process(nums, 0, target, 0, &res)
//     return res
// }
//
// func process(nums []int, index, target, cur int, res *int) {
//     if index == len(nums) {
//         if cur == target {
//             *res++
//         }
//         return
//     }
//     process(nums, index+1, target, cur+nums[index], res)
//     process(nums, index+1, target, cur-nums[index], res)
// }

// 动态规划
func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, v := range nums {
        sum+=v
    }
    if math.Abs(float64(target)) > math.Abs(float64(sum)) {
        return 0
    }
    n := len(nums)
    // - 0 +
    t := 2*sum+1
    dp := make([][]int, n)
    for i:=0;i<len(dp);i++ {
        dp[i] = make([]int, t)
    }
    if nums[0] == 0 {
        // +0 -0都得到0
        dp[0][sum] = 2
    } else {
        dp[0][sum + nums[0]] = 1
        dp[0][sum - nums[0]] = 1
    }

    for i:=1;i<n;i++ {
        for j:=0;j<t;j++ {
            l, r := 0, 0
            if j - nums[i] >= 0 {
                l = j-nums[i]
            }
            if j + nums[i] < t {
                r = j + nums[i]
            }
            dp[i][j] = dp[i-1][l] + dp[i-1][r]
        }
    }
    return dp[n-1][sum+target]
}
//leetcode submit region end(Prohibit modification and deletion)
