//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
//字，并以数组的形式返回结果。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,1]
//输出：[2]
// 
//
// 
//
// 提示： 
//
// 
// n == nums.length 
// 1 <= n <= 10⁵ 
// 1 <= nums[i] <= n 
// 
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。 
//
// Related Topics 数组 哈希表 👍 1323 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
    // -------------------------- O(n) 时间复杂度 O(1) 空间复杂度 --------------------------
    // 由于 nums[i] 在 [1, n] 内
    // 1. 遍历 nums, 将 nums[nums[i] - 1] 取反
    // 2. 再次遍历 nums, 若 nums[i] 为正数, 则 i + 1 未出现
    ans := []int{}
    for i := 0; i < len(nums); i++ {
        // 取反
        if nums[abs(nums[i])-1] > 0 {
            nums[abs(nums[i])-1] *= -1
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            ans = append(ans, i+1)
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
//leetcode submit region end(Prohibit modification and deletion)
