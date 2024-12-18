//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 200 
// 1 <= nums[i] <= 100 
// 
//
// Related Topics 数组 动态规划 👍 2192 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func canPartition(nums []int) bool {
//     if len(nums) < 2 {
//         return false
//     }
//     sum, maxNum := 0, 0
//     for _, v := range nums {
//         if maxNum < v {
//             maxNum = v
//         }
//         sum+=v
//     }
//     if sum%2 != 0 {
//         return false
//     }
//     target := sum/2
//     if maxNum > target {
//         return false
//     }
//     dp := make([][]bool, len(nums))
//     for i:=0;i<len(dp);i++ {
//         dp[i] = make([]bool, target+1)
//         dp[i][0] = true
//     }
//     dp[0][nums[0]] = true
//     for i:=1;i<len(dp);i++ {
//         v := nums[i]
//         for j:=1;j<=target;j++ {
//             if j >= v {
//                 dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
//             } else {
//                 dp[i][j] = dp[i-1][j]
//             }
//         }
//     }
//     return dp[len(dp)-1][target]
// }

// 动态规划降维
// func canPartition(nums []int) bool {
// n := len(nums)
// if n < 2 {
// return false
// }
//
// sum, max := 0, 0
// for _, v := range nums {
// sum += v
// if v > max {
// max = v
// }
// }
// if sum%2 != 0 {
// return false
// }
//
// target := sum / 2
// if max > target {
// return false
// }
//
// dp := make([]bool, target+1)
// dp[0] = true
// for i := 0; i < n; i++ {
// v := nums[i]
// for j := target; j >= v; j-- {
// dp[j] = dp[j] || dp[j-v]
// }
// }
// return dp[target]
// }

func canPartition(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }
    sum, max := 0, 0
    for _, v := range nums {
        sum += v
        if v > max {
            max = v
        }
    }
    if sum%2 != 0 {
        return false
    }
    target := sum/2
    if max > target {
        return false
    }

    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, target+1)
    }
    for i:=0;i<n;i++ {
        dp[i][0] = true
    }
    dp[0][nums[0]] = true
    for i:=1;i<n;i++ {
        v := nums[i]
        for j := 1;j<=target;j++ {
            if j>=v {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)
