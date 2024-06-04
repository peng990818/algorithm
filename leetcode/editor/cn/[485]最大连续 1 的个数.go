//给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,0,1,1,1]
//输出：3
//解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
// 
//
// 示例 2: 
//
// 
//输入：nums = [1,0,1,1,0,1]
//输出：2
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// nums[i] 不是 0 就是 1. 
// 
//
// Related Topics 数组 👍 436 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
    res, tmp := 0, 0
    for _, v := range nums {
        if v == 1 {
            tmp++
        } else {
            if res < tmp {
                res = tmp
            }
            tmp = 0
        }
    }
    if res < tmp {
        res = tmp
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
