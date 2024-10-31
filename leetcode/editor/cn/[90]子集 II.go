//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。 
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。 
//
// 
// 
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
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
// 
//
// Related Topics 位运算 数组 回溯 👍 1245 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// flag 用于标记选没选过前一个数
func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    path := make([]int, 0)
    process(&res, nums, path, 0, false)
    return res
}

func process(res *[][]int, nums []int, path []int, cur int, flag bool) {
    tmp := make([]int, len(path))
    copy(tmp, path)
    *res = append(*res, tmp)
    for i:=cur;i<len(nums);i++ {
        if !flag && i>0 && nums[i-1] == nums[i] {
            continue
        }
        flag = true
        path = append(path, nums[i])
        tmp := i
        tmp++
        process(res, nums, path, tmp, flag)
        path = path[:len(path)-1]
        flag = false
    }
}
//leetcode submit region end(Prohibit modification and deletion)
