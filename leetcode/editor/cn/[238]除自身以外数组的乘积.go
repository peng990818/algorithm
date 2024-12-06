//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。 
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。 
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
// 
//
// 示例 2: 
//
// 
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]
// 
//
// 
//
// 提示： 
//
// 
// 2 <= nums.length <= 10⁵ 
// -30 <= nums[i] <= 30 
// 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内 
// 
//
// 
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。） 
//
// Related Topics 数组 前缀和 👍 1895 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func productExceptSelf(nums []int) []int {
//     answer := make([]int, len(nums))
//     // 申请空间存储数组中每个点的左边乘积和右边乘积
//     L, R := make([]int, len(nums)), make([]int, len(nums))
//     L[0] = 1
//     for i:=1;i<len(nums);i++ {
//         L[i] = L[i-1]*nums[i-1]
//     }
//
//     R[len(nums)-1] = 1
//     for i:=len(nums)-2;i>=0;i-- {
//         R[i] = R[i+1]*nums[i+1]
//     }
//
//     for i:=0;i<len(nums);i++ {
//         answer[i] = L[i] * R[i]
//     }
//     return answer
// }

// func productExceptSelf(nums []int) []int {
//     answer := make([]int, len(nums))
//     answer[0] = 1
//     for i:=1;i<len(nums);i++ {
//         answer[i] = answer[i-1]*nums[i-1]
//     }
//
//     R := 1
//     for i:=len(nums)-1;i>=0;i-- {
//         answer[i] *= R
//         R *= nums[i]
//     }
//     return answer
// }

// 申请空间
// func productExceptSelf(nums []int) []int {
//     if len(nums) == 0 {
//         return nil
//     }
//     L := make([]int, len(nums))
//     R := make([]int, len(nums))
//
//     L[0] = 1
//     for i:=1;i<len(nums);i++ {
//         L[i] = L[i-1] * nums[i-1]
//     }
//
//     R[len(nums)-1] = 1
//     for i:=len(nums)-2;i>=0;i-- {
//         R[i] = R[i+1] * nums[i+1]
//     }
//
//     answer := make([]int, len(nums))
//     for i:=0;i<len(nums);i++ {
//         answer[i] = L[i] * R[i]
//     }
//     return answer
// }

// 不申请空间
func productExceptSelf(nums []int) []int {
    if len(nums) == 0 {
        return nil
    }
    answers := make([]int, len(nums))
    answers[0] = 1
    for i:=1;i<len(nums);i++ {
        answers[i] = nums[i-1] * answers[i-1]
    }

    R := 1
    for i:=len(answers)-1;i>=0;i-- {
        answers[i] *= R
        R *= nums[i]
    }
    return answers
}
//leetcode submit region end(Prohibit modification and deletion)
