//给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。 
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// 0 <= nums[i] <= 10⁵ 
// 
//
// Related Topics 贪心 数组 动态规划 👍 2890 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func canJump(nums []int) bool {
//     if len(nums) == 0 {
//         return false
//     }
//     if len(nums) == 1 && nums[0] == 0 {
//         return true
//     }
//     i := 0
//     for i < len(nums) {
//         if nums[i] == 0 {
//             return false
//         }
//         // 一步到位
//         if i+nums[i] >= len(nums)-1 {
//             return true
//         }
//         // 贪心寻找最远的点
//         j, max, index := i+1, nums[i] + nums[i+nums[i]], i+nums[i]
//         for j < len(nums) && j < i+nums[i] {
//             tmp := j-i+nums[j]
//             if tmp > max {
//                 max = tmp
//                 index = j
//             }
//             j++
//         }
//         i = index
//     }
//     return true
// }

// 贪心策略：计算每一次跳的最远距离
// func canJump(nums []int) bool {
//     if len(nums) == 0 {
//         return false
//     }
//     if len(nums) == 1 && nums[0] == 0 {
//         return true
//     }
//     i := 0
//     for i<len(nums) {
//         // 零值特殊处理
//         if nums[i] == 0 {
//             return false
//         }
//         // 一步到位
//         if i + nums[i] >= len(nums)-1 {
//             return true
//         }
//         // 计算最远距离
//         max := nums[i] + nums[i+nums[i]]
//         j := i+1
//         index := i+nums[i]
//         for j < len(nums) && j < i+nums[i] {
//             tmp := nums[j] + j-i
//             if tmp > max {
//                 max = tmp
//                 index = j
//             }
//             j++
//         }
//         i = index
//     }
//     return true
// }

// 真正的贪心
// func canJump(nums []int) bool {
//     n := len(nums)
//     most := 0
//     for i:=0;i<n;i++ {
//         if i <= most {
//             // 位置都可达
//             most = max(most, i+nums[i])
//             if most >= n-1 {
//                 return true
//             }
//         } else {
//             return false
//         }
//     }
//     return false
// }

func canJump(nums []int) bool {
if len(nums) == 0 {
return false
}
if len(nums) == 1 {
return true
}
i := 0
end := nums[i] + i
for i <= end && i < len(nums) {
if end >= len(nums)-1 {
return true
}
end = max(nums[i] + i, end)
i++
}
return false
}









//leetcode submit region end(Prohibit modification and deletion)
