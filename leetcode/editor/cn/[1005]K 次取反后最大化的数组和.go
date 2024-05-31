//给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组： 
//
// 
// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。 
// 
//
// 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。 
//
// 以这种方式修改数组后，返回数组 可能的最大和 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [4,2,3], k = 1
//输出：5
//解释：选择下标 1 ，nums 变为 [4,-2,3] 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,-1,0,2], k = 3
//输出：6
//解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
// 
//
// 示例 3： 
//
// 
//输入：nums = [2,-3,-1,5,-4], k = 2
//输出：13
//解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// -100 <= nums[i] <= 100 
// 1 <= k <= 10⁴ 
// 
//
// Related Topics 贪心 数组 排序 👍 450 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMin(nums []int) int {
    n := len(nums)
    min := nums[0]
    minIndex := 0
    for i:=1;i<n;i++ {
        if nums[i] < min {
            min = nums[i]
            minIndex = i
        }
    }
    return minIndex
}

func largestSumAfterKNegations(nums []int, k int) int {
    for i:=0;i<k;i++ {
        index := findMin(nums)
        if nums[index] == 0 {
            break
        } else {
            nums[index]*=-1
        }
    }
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return sum
}
//leetcode submit region end(Prohibit modification and deletion)
