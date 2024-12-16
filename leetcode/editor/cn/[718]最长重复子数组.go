//给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//输出：3
//解释：长度最长的公共子数组是 [3,2,1] 。
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
//输出：5
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length, nums2.length <= 1000 
// 0 <= nums1[i], nums2[i] <= 100 
// 
//
// Related Topics 数组 二分查找 动态规划 滑动窗口 哈希函数 滚动哈希 👍 1136 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums1, nums2 []int, i, j int) int {
    size1, size2  := len(nums1), len(nums2)
    max_len, cur_len := 0, 0
    for i<size1 && j<size2 {
        if nums1[i] == nums2[j] {
            cur_len++
            max_len = max(max_len, cur_len)
        }
        i++
        j++
    }
    return max_len
}

func findLength(nums1 []int, nums2 []int) int {
    res := 0
    for i := range nums1 {
        res = max(res, findMaxLength(nums1, nums2, i, 0))
    }
    for i:=range nums2 {
        res = max(res, findMaxLength(nums1, nums2, 0, i))
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
