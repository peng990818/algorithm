//给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//解释：[4,9] 也是可通过的
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length, nums2.length <= 1000 
// 0 <= nums1[i], nums2[i] <= 1000 
// 
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 946 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 哈希表
// func intersection(nums1 []int, nums2 []int) []int {
//     mp1 := make(map[int]struct{}, len(nums1))
//     mp2 := make(map[int]struct{}, len(nums2))
//     for _, v := range nums1 {
//         mp1[v] = struct{}{}
//     }
//
//     res := make([]int, 0)
//     for _, v := range nums2 {
//         mp2[v] = struct{}{}
//     }
//
//     for k, _ := range mp1 {
//         if _, ok := mp2[k]; ok {
//             res = append(res, k)
//         }
//     }
//     return res
// }

// 排序+双指针
func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res := make([]int, 0, len(nums2))
    i, j := 0, 0
    for i<len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            if len(res) == 0 || nums1[i] > res[len(res)-1] {
                res = append(res, nums1[i])
            }
            i++
            j++
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
