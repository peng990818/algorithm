//给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。 
//
// 请你找出符合题意的 最短 子数组，并输出它的长度。 
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
//输入：nums = [2,6,4,8,10,9,15]
//输出：5
//解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,4]
//输出：0
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// -10⁵ <= nums[i] <= 10⁵ 
// 
//
// 
//
// 进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？ 
//
// Related Topics 栈 贪心 数组 双指针 排序 单调栈 👍 1178 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findUnsortedSubarray(nums []int) int {
    // if sort.IntsAreSorted(nums) {
    //     return 0
    // }
    // 排序比较
    // numsSorted := append([]int(nil), nums...)
    // sort.Ints(numsSorted)
    // left, right := 0, len(nums)-1
    // for nums[left] == numsSorted[left] {
    //     left++
    // }
    // for nums[right] == numsSorted[right] {
    //     right--
    // }
    // return right-left+1

    n := len(nums)
    minn, maxn := math.MaxInt64, math.MinInt64
    left, right := -1, -1
    for i, num := range nums {
        if maxn > num {
            right = i
        } else {
            maxn = num
        }
        if minn < nums[n-i-1] {
            left = n-i-1
        } else {
            minn = nums[n-i-1]
        }
    }
    if right == -1 {
        return 0
    }
    return right-left+1
}
//leetcode submit region end(Prohibit modification and deletion)
