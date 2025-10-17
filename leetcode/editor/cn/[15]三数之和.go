//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != 
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
// 
//
// 示例 3： 
//
// 
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 3000 
// -10⁵ <= nums[i] <= 10⁵ 
// 
//
// Related Topics 数组 双指针 排序 👍 7147 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func threeSum(nums []int) [][]int {
//     if len(nums) <= 2 {
//         return nil
//     }
//     res := make([][]int, 0)
//     sort.Ints(nums)
//     n := len(nums)
//     for i:=0;i<n;i++ {
//         if i != 0 && nums[i-1] == nums[i] {
//             continue
//         }
//         sum := 0-nums[i]
//         l, r := i+1, len(nums)-1
//         for l<r {
//             if l > i+1 && nums[l] == nums[l-1] {
//                 l++
//                 continue
//             }
//             tmp := nums[l] + nums[r]
//             if tmp < sum {
//                 l++
//             } else if tmp > sum {
//                 r--
//             } else {
//                 res = append(res, []int{nums[i], nums[l], nums[r]})
//                 l++
//             }
//         }
//     }
//     return res
// }

// func threeSum(nums []int) [][]int {
//     if len(nums) <= 2 {
//         return nil
//     }
//     sort.Ints(nums)
//     res := make([][]int, 0)
//     for i:=0;i<len(nums)-2;i++ {
//         // 去重
//         if i != 0 && nums[i-1] == nums[i] {
//             continue
//         }
//         sum := -nums[i]
//         l, r := i+1, len(nums)-1
//         for l < r {
//             // 去重
//             if l > i+1 && nums[l-1] == nums[l] {
//                 l++
//                 continue
//             }
//             tmp := nums[l] + nums[r]
//             if tmp < sum {
//                 l++
//             } else if tmp > sum {
//                 r--
//             } else {
//                 res = append(res, []int{nums[i], nums[l], nums[r]})
//                 l++
//             }
//         }
//     }
//     return res
// }

// func threeSum(nums []int) [][]int {
// sort.Ints(nums)
// var res [][]int
// for i:=0;i<len(nums)-2;i++ {
// if nums[i] > 0 {
// continue
// }
// if i>0 && nums[i] == nums[i-1] {
// continue
// }
// l, r := i+1, len(nums)-1
// for l < r {
// if nums[i] + nums[l] + nums[r] < 0 {
// l++
// } else if nums[i] + nums[l] + nums[r] > 0 {
// r--
// } else {
// res = append(res, []int{nums[i], nums[l], nums[r]})
// for r > l && nums[r] == nums[r-1] {
// r--
// }
// for l < r && nums[l] == nums[l+1] {
// l++
// }
// l++
// r--
// }
// }
// }
// return res
// }

func threeSum(nums []int) [][]int {
var res [][]int
sort.Ints(nums)
for i:=0;i<len(nums);i++ {
if nums[i] > 0 {
continue
}
if i > 0 && nums[i] == nums[i-1] {
continue
}
l, r := i+1, len(nums)-1
for l < r {
if nums[i] + nums[l] + nums[r] == 0 {
res = append(res, []int{nums[i], nums[l], nums[r]})
for r > l && nums[r] == nums[r-1] {
r--
}
for l < r && nums[l] == nums[l+1] {
l++
}
l++
r--
} else if nums[i] + nums[l] + nums[r] < 0 {
l++
} else {
r--
}
}

}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
