//给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。 
//
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。 
//
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,4,2,2]
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,1,3,4,2]
//输出：3
// 
//
// 示例 3 : 
//
// 
//输入：nums = [3,3,3,3,3]
//输出：3
// 
//
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 10⁵ 
// nums.length == n + 1 
// 1 <= nums[i] <= n 
// nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次 
// 
//
// 
//
// 进阶： 
//
// 
// 如何证明 nums 中至少存在一个重复的数字? 
// 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？ 
// 
//
// Related Topics 位运算 数组 双指针 二分查找 👍 2500 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func findDuplicate(nums []int) int {
//     // 双指针
//     slow, fast := nums[0], nums[nums[0]]
//     for slow != fast {
//         slow = nums[slow]
//         fast = nums[nums[fast]]
//     }
//     help := 0
//     for help != slow {
//         help = nums[help]
//         slow = nums[slow]
//     }
//     return slow
// }

// 二分法
// func findDuplicate(nums []int) int {
//     l, r := 1, len(nums)
//     for l < r {
//         mid := (l+r) / 2
//         cnt := 0
//         for _, v := range nums {
//             if v >= l && v <= mid {
//                 cnt++
//             }
//         }
//         if cnt > mid-l+1 {
//             r = mid
//         } else {
//             l = mid+1
//         }
//     }
//     return l
// }

// 快慢指针
// func findDuplicate(nums []int) int {
//     slow, fast := 0, 0
//     for {
//         fast = nums[fast]
//         fast = nums[fast]
//         slow = nums[slow]
//         if slow == fast {
//             break
//         }
//     }
//
//     fast = 0
//     for fast != slow {
//         fast = nums[fast]
//         slow = nums[slow]
//     }
//     return fast
// }

// 二分法
// func findDuplicate(nums []int) int {
//     min, max := 1, len(nums)
//     for min < max {
//         mid := (min+max) / 2
//         cnt := 0
//         for _, v := range nums {
//             if v >= min && v <= mid {
//                 cnt++
//             }
//         }
//         if cnt > mid - min + 1 {
//             max = mid
//         } else {
//             min = mid+1
//         }
//     }
//     return min
// }

// 双指针法
func findDuplicate(nums []int) int {
    slow, fast := 0, 0
    for {
        // fast 前进两次， slow 前进一次， 相等证明存在环
        fast = nums[fast]
        fast = nums[fast]
        slow = nums[slow]
        if slow == fast {
            break
        }
    }

    // 从开始和相交处开始遍历，相等时即为环的入口，也就是重复元素
    p := 0
    for p != slow {
        p = nums[p]
        slow = nums[slow]
    }
    return p
}
//leetcode submit region end(Prohibit modification and deletion)
