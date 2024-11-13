//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。 
//
// 子数组是数组中元素的连续非空序列。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,1], k = 2
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3], k = 3
//输出：2
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2 * 10⁴ 
// -1000 <= nums[i] <= 1000 
// -10⁷ <= k <= 10⁷ 
// 
//
// Related Topics 数组 哈希表 前缀和 👍 2531 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 优化
// func subarraySum(nums []int, k int) int {
//     count := 0
//     for i:=0;i<len(nums);i++ {
//         sum := 0
//         for j:=i;j<len(nums);j++ {
//             sum+=nums[j]
//             if sum == k {
//                 count++
//             }
//         }
//     }
//     return count
// }


func subarraySum(nums []int, k int) int {
    count, pre := 0, 0
    m := map[int]int{}
    m[0] = 1
    for i:=0;i<len(nums);i++ {
        pre += nums[i]
        if _, ok := m[pre-k];ok{
            count += m[pre-k]
        }
        m[pre] += 1
    }
    return count
}
//leetcode submit region end(Prohibit modification and deletion)
