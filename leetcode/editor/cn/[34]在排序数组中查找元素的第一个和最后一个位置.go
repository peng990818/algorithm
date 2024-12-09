//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。 
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。 
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4] 
//
// 示例 2： 
//
// 
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1] 
//
// 示例 3： 
//
// 
//输入：nums = [], target = 0
//输出：[-1,-1] 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 10⁵ 
// -10⁹ <= nums[i] <= 10⁹ 
// nums 是一个非递减数组 
// -10⁹ <= target <= 10⁹ 
// 
//
// Related Topics 数组 二分查找 👍 2842 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func searchRange(nums []int, target int) []int {
//     if len(nums) == 0 {
//         return []int{-1, -1}
//     }
//     // 找到数组中第一个等于target的位置，作为左边界
//     left := search(nums, target, true)
//     // 找到数组中第一个大于target的位置，作为右边界
//     right := search(nums, target, false) - 1;
//     if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
//         return []int{left, right}
//     }
//     return []int{-1, -1}
// }
//
// func search(nums []int, target int, lower bool) int {
//     l, r := 0, len(nums)-1
//     res := len(nums)
//     for l <= r {
//         mid := l + ((r-l)>>1)
//         if (nums[mid] > target || (lower && nums[mid] >= target)) {
//             r = mid-1
//             res = mid
//         } else {
//             l = mid+1
//         }
//     }
//     return res
// }

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    // 找到第一个等于target的索引，视为左边界
    left := binarySearch(nums, target, true)
    // 找到第一个大于target的索引，再减去1，视为右边界
    right := binarySearch(nums, target, false) - 1
    if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
        return []int{left, right}
    }
    return []int{-1, -1}
}

func binarySearch(nums []int, target int, flag bool) int {
    l, r := 0, len(nums)-1
    res := len(nums)
    for l<=r {
        mid := l + (r-l)/2
        if nums[mid] > target || (flag && nums[mid] >= target) {
            r = mid-1
            res = mid
        } else {
            l = mid+1
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
