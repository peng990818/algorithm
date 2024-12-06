//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [3,2,3]
//输出：3 
//
// 示例 2： 
//
// 
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
// 
//
// 
//提示：
//
// 
// n == nums.length 
// 1 <= n <= 5 * 10⁴ 
// -10⁹ <= nums[i] <= 10⁹ 
// 
//
// 
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。 
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 2340 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 哈希表
// func majorityElement(nums []int) int {
//     mp := make(map[int]int)
//     for _, num := range nums {
//         mp[num]++
//     }
//     for k, v := range mp {
//         if v > len(nums)/2 {
//             return k
//         }
//     }
//     return 0
// }

// 不申请空间
// func majorityElement(nums []int) int {
//     sort.Ints(nums)
//     cnt := 1
//     for i:=1;i<len(nums);i++ {
//         if nums[i] == nums[i-1] {
//             cnt++
//         } else {
//             if cnt > len(nums)/2 {
//                 return nums[i-1]
//             }
//             cnt = 1
//         }
//     }
//     if cnt > len(nums)/2 {
//         return nums[len(nums)-1]
//     }
//     return 0
// }

func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}
//leetcode submit region end(Prohibit modification and deletion)
