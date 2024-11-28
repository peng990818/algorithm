//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。 
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0]
//输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// nums 中的所有元素 互不相同 
// 
//
// Related Topics 位运算 数组 回溯 👍 2381 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func subsets(nums []int) [][]int {
//     res := make([][]int, 0)
//     set := make([]int, 0)
//     process(&res, set, nums, 0)
//     return res
// }
//
// func process(res *[][]int, set []int, nums []int, cur int) {
//     tmp := make([]int, len(set))
//     copy(tmp, set)
//     *res = append(*res, tmp)
//     for i:=cur;i<len(nums);i++ {
//         // 选择这个数
//         set = append(set, nums[cur])
//         cur++
//         process(res, set, nums, cur)
//         // 不选择这个数
//         set = set[:len(set)-1]
//     }
// }


func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    var process func (path []int, index int)
    process = func (path []int, index int) {
        if index == len(nums) {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        path = append(path, nums[index])
        process(path, index+1)
        path = path[:len(path)-1]
        process(path, index+1)
    }
    process(nil, 0)
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
