//给定一个未排序的整数数组
// nums ， 返回最长递增子序列的个数 。 
//
// 注意 这个数列必须是 严格 递增的。 
//
// 
//
// 示例 1: 
//
// 
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
// 
//
// 示例 2: 
//
// 
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
// 
//
// 
//
// 提示: 
//
// 
// 
//
// 
// 1 <= nums.length <= 2000 
// -10⁶ <= nums[i] <= 10⁶ 
// 
//
// Related Topics 树状数组 线段树 数组 动态规划 👍 957 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findNumberOfLIS(nums []int) int {
if len(nums) <= 1 {
return len(nums)
}
maxCount := 0
dp := make([]int, len(nums))
for i:=range dp {
dp[i] = 1
}
count := make([]int, len(nums))
for i:=range count {
count[i] = 1
}
for i:=1;i<len(nums);i++ {
for j:=0;j<i;j++ {
if nums[i] > nums[j] {
if dp[j]+1 > dp[i] {
dp[i] = dp[j]+1
count[i] = count[j]
} else if dp[i] == dp[j]+1 {
count[i] += count[j]
}
}
maxCount = max(maxCount, dp[i])
}
}
res := 0
for i:=0;i<len(nums);i++ {
if maxCount == dp[i] {
res += count[i]
}
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
