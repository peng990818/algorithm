//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
// 
//
// 示例 2: 
//
// 
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释: 
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// -2³¹ <= nums[i] <= 2³¹ - 1 
// 0 <= k <= 10⁵ 
// 
//
// 
//
// 进阶： 
//
// 
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。 
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？ 
// 
//
// Related Topics 数组 数学 双指针 👍 2425 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func rotate(nums []int, k int)  {
// k = k % len(nums)
// for i := 0;i<k;i++ {
// tmp := nums[len(nums)-1]
// for j := len(nums)-1;j>0;j-- {
// nums[j] = nums[j-1]
// }
// nums[0] = tmp
// }
// }

func rotate(nums []int, k int)  {
// 三次旋转
k = k % len(nums)
// 先反转这个字符串
for i,j:=0,len(nums)-1;i<j;i, j = i+1, j-1 {
nums[i], nums[j] = nums[j], nums[i]
}
// 再反转前k个字符串
for i,j:=0,k-1;i<j;i,j=i+1,j-1 {
nums[i], nums[j] = nums[j], nums[i]
}
// 最后反转k到末尾
for i,j:=k, len(nums)-1;i<j;i,j=i+1,j-1 {
nums[i], nums[j] = nums[j], nums[i]
}
}
//leetcode submit region end(Prohibit modification and deletion)
