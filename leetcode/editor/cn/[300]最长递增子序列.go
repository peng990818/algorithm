//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。 
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。 
//
// 示例 1： 
//
// 
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,0,3,2,3]
//输出：4
// 
//
// 示例 3： 
//
// 
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2500 
// -10⁴ <= nums[i] <= 10⁴ 
// 
//
// 
//
// 进阶： 
//
// 
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗? 
// 
//
// Related Topics 数组 二分查找 动态规划 👍 3787 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 1、动态规划
// func lengthOfLIS(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }
//     dp := make([]int, len(nums))
//     maxNum := 0
//     for i:=0;i<len(nums);i++ {
//         dp[i] = 1
//         for j:=0;j<i;j++ {
//             if nums[j] < nums[i] {
//                 dp[i] = max(dp[i], dp[j]+1)
//             }
//         }
//         if maxNum < dp[i] {
//             maxNum = dp[i]
//         }
//     }
//     return maxNum
// }

// 2、贪心+二分
// func lengthOfLIS(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }
//     d := make([]int, 1)
//     d[0] = nums[0]
//     for i:=1;i<len(nums);i++ {
//         if nums[i] > d[len(d)-1] {
//             d = append(d, nums[i])
//         } else {
//             // 二分
//             l, r := 0, len(d)-1
//             pos := r
//             for l <= r {
//                 mid := (l+r) >> 1
//                 if d[mid] < nums[i] {
//                     l = mid + 1
//                 } else {
//                     pos = mid
//                     r = mid - 1
//                 }
//             }
//             d[pos] = nums[i]
//         }
//     }
//     return len(d)
// }

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, len(nums))
    maxNum := 0
    for i:=0;i<len(nums);i++ {
        dp[i] = 1
        for j:=0;j<i;j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        maxNum = max(maxNum, dp[i])
    }
    return maxNum
}

//leetcode submit region end(Prohibit modification and deletion)
