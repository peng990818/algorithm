//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 8 
// -10 <= nums[i] <= 10 
// 
//
// Related Topics 数组 回溯 👍 1617 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func process(res *[][]int, arr []int, first int) {
    if first == len(arr) {
        tmp := make([]int, len(arr))
        copy(tmp, arr)
        *res = append(*res, tmp)
    }

    mp := make(map[int]bool)
    for i:=first;i<len(arr);i++ {
        if mp[arr[i]] {
            continue
        }
        arr[i], arr[first] = arr[first], arr[i]
        process(res, arr, first+1)
        arr[i], arr[first] = arr[first], arr[i]
        mp[arr[i]] = true
    }
}

func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    process(&res, nums, 0)
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
