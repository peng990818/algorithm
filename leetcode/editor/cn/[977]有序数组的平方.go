//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。 
//
// 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100] 
//
// 示例 2： 
//
// 
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// -10⁴ <= nums[i] <= 10⁴ 
// nums 已按 非递减顺序 排序 
// 
//
// 
//
// 进阶： 
//
// 
// 请你设计时间复杂度为 O(n) 的算法解决本问题 
// 
//
// Related Topics 数组 双指针 排序 👍 1113 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 暴力排序
// func sortedSquares(nums []int) []int {
// if len(nums) == 0 {
// return nil
// }
//
// for i := 0; i < len(nums); i++ {
// nums[i] = nums[i] * nums[i]
// }
// sort.Ints(nums)
// return nums
// }

// func sortedSquares(nums []int) []int {
// if len(nums) == 0 {
// return nil
// }
//
// res := make([]int, len(nums))
// i, j, k := 0, len(nums)-1, len(nums)-1
// for i <= j {
// tI := nums[i] * nums[i]
// tJ := nums[j] * nums[j]
// if tI >= tJ {
// res[k] = nums[i] * nums[i]
// i++
// } else {
// res[k] = nums[j] * nums[j]
// j--
// }
// k--
// }
// return res
// }

// func sortedSquares(nums []int) []int {
// res := make([]int, len(nums))
// k := len(nums) - 1
// for i, j := 0, len(nums)-1;i<=j; {
// if nums[i] * nums[i] < nums[j] * nums[j] {
// res[k] = nums[j]*nums[j]
// j--
// } else {
// res[k] = nums[i]*nums[i]
// i++
// }
// k--
// }
// return res
// }

// func sortedSquares(nums []int) []int {
// sort.Slice(nums, func(i, j int) bool {
// return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
// })
// for i:=0;i<len(nums);i++ {
// nums[i] *= nums[i]
// }
// return nums
// }

func sortedSquares(nums []int) []int {
res := make([]int, len(nums))
k := len(nums)-1
for i,j:=0,len(nums)-1;i<=j;{
if nums[i]*nums[i] < nums[j]*nums[j] {
res[k] = nums[j]*nums[j]
j--
} else {
res[k] = nums[i]*nums[i]
i++
}
k--
}
return res
}

//leetcode submit region end(Prohibit modification and deletion)
