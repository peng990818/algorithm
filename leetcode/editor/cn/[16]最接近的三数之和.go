//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。 
//
// 返回这三个数的和。 
//
// 假定每组输入只存在恰好一个解。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,0,0], target = 1
//输出：0
//解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 1000 
// -1000 <= nums[i] <= 1000 
// -10⁴ <= target <= 10⁴ 
// 
//
// Related Topics 数组 双指针 排序 👍 1680 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
    if len(nums) <= 2 {
        return 0
    }
    sort.Ints(nums)
    res, c := math.MaxInt32, math.MaxInt32
    for i:=0;i<len(nums)-2;i++ {
        l, r := i+1, len(nums)-1
        for l < r {
            tmp := nums[l] + nums[r] + nums[i]
            if tmp <= target {
                if target - tmp < c {
                    c = target - tmp
                    res = tmp
                }
                l++
            } else {
                if tmp - target < c {
                    c = tmp - target
                    res = tmp
                }
                r--
            }
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
