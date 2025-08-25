//给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。 
//
// 题目数据保证答案符合 32 位整数范围。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3], target = 4
//输出：7
//解释：
//所有可能的组合为：
//(1, 1, 1, 1)
//(1, 1, 2)
//(1, 2, 1)
//(1, 3)
//(2, 1, 1)
//(2, 2)
//(3, 1)
//请注意，顺序不同的序列被视作不同的组合。
// 
//
// 示例 2： 
//
// 
//输入：nums = [9], target = 3
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 200 
// 1 <= nums[i] <= 1000 
// nums 中的所有元素 互不相同 
// 1 <= target <= 1000 
// 
//
// 
//
// 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？ 
//
// Related Topics 数组 动态规划 👍 1162 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 二维
// 排列的情况下外层循环最好是target，不用限制使用的顺序
// func combinationSum4(nums []int, target int) int {
// n := len(nums)
// dp := make([][]int, target+1)
// for i := range dp {
// dp[i] = make([]int, n+1)
// }
//
// dp[0][0] = 1
//
// for i:=0;i<=target;i++ {
// for j:=1;j<=n;j++ {
// dp[i][j] = dp[i][j-1]
// if i >= nums[j-1] {
// // 用 nums[j-1]，顺序敏感 -> 从dp[i-nums[j-1]][n] 转移
// dp[i][j] += dp[i-nums[j-1]][n]
// }
// }
// }
// return dp[target][n]
// }

// 一维
// func combinationSum4(nums []int, target int) int {
// dp := make([]int, target+1)
// // 所有数字都不用
// dp[0] = 1
// for i:=1;i<=target;i++ {
// for j:=0;j<len(nums);j++ {
// if i >= nums[j] {
// dp[i] += dp[i-nums[j]]
// }
// }
// }
// return dp[target]
// }

func combinationSum4(nums []int, target int) int {
dp := make([]int, target+1)
dp[0] = 1
n := len(nums)
// 注意顺序，属于排列问题
for j:=1;j<=target;j++ {
for i:=0;i<n;i++ {
if j >= nums[i] {
dp[j] += dp[j-nums[i]]
}
}
}
return dp[target]
}
//leetcode submit region end(Prohibit modification and deletion)
