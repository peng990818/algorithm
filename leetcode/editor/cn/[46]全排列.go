//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 6 
// -10 <= nums[i] <= 10 
// nums 中的所有整数 互不相同 
// 
//
// Related Topics 数组 回溯 👍 2905 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func process(res *[][]int, arr []int, first int) {
//     if first == len(arr) {
//         tmp := make([]int, len(arr))
//         copy(tmp, arr)
//         *res = append(*res, tmp)
//     }
//
//     // 轮流让数组中每个数都做一次第一个数，递归让后面的数全排列
//     for i:=first;i<len(arr);i++ {
//         arr[i], arr[first] = arr[first], arr[i]
//         process(res, arr, first+1)
//         arr[i], arr[first] = arr[first], arr[i]
//     }
// }
//
//
// func permute(nums []int) [][]int {
//     res := make([][]int, 0)
//     process(&res, nums, 0)
//     return res
// }

func permute(nums []int) [][]int {
    res := make([][]int, 0)
    process(&res, nums, 0)
    return res
}

func process(res *[][]int, nums []int, index int) {
    if index == len(nums) {
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        *res = append(*res, tmp)
    }
    for i:=index;i<len(nums);i++ {
        nums[i], nums[index] = nums[index], nums[i]
        process(res, nums, index+1)
        nums[i], nums[index] = nums[index], nums[i]
    }
}

//leetcode submit region end(Prohibit modification and deletion)
