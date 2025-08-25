//nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。 
//
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。 
//
// 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 
//nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。 
//
// 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出：[-1,3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
//- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
//- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。 
//
// 示例 2： 
//
// 
//输入：nums1 = [2,4], nums2 = [1,2,3,4].
//输出：[3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
//- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length <= nums2.length <= 1000 
// 0 <= nums1[i], nums2[i] <= 10⁴ 
// nums1和nums2中所有整数 互不相同 
// nums1 中的所有整数同样出现在 nums2 中 
// 
//
// 
//
// 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？ 
//
// Related Topics 栈 数组 哈希表 单调栈 👍 1177 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
//     mp := make(map[int]int, len(nums1))
//     res := make([]int, len(nums1))
//     stack := make([]int, 0)
//     for j:=len(nums2)-1;j>=0;j-- {
//         num := nums2[j]
//         for len(stack) > 0 && num >= stack[len(stack)-1] {
//             stack = stack[:len(stack)-1]
//         }
//         if len(stack) > 0 {
//             mp[num] = stack[len(stack)-1]
//         } else {
//             mp[num] = -1
//         }
//         stack = append(stack, num)
//     }
//
//     for i, num := range nums1 {
//         res[i] = mp[num]
//     }
//     return res
// }
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// tmp := make([]int, len(nums2))
// for i:=range tmp {
// tmp[i] = -1
// }
// stack := make([]int, 0)
// for i:=0;i<len(nums2);i++ {
// for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
// top := stack[len(stack)-1]
// tmp[top] = i
// stack = stack[:len(stack)-1]
// }
// stack = append(stack, i)
// }
// res := make([]int, len(nums1))
// for i := range nums1 {
// for j:=range nums2 {
// if nums1[i] == nums2[j] {
// if tmp[j] >= 0 {
// res[i] = nums2[tmp[j]]
// break
// }
// res[i] = tmp[j]
// }
// }
// }
// return res
// }

func nextGreaterElement(nums1 []int, nums2 []int) []int {
mp := make(map[int]int, len(nums1))
for i:=range nums1 {
mp[nums1[i]] = -1
}

stack := make([]int, 0)
for i:=0;i<len(nums2);i++ {
for len(stack)>0 && nums2[i] > stack[len(stack)-1] {
top := stack[len(stack)-1]
if _, ok := mp[top];ok {
mp[top] = nums2[i]
}
stack = stack[:len(stack)-1]
}
stack = append(stack, nums2[i])
}

res := make([]int, len(nums1))
for i :=range nums1 {
res[i] = mp[nums1[i]]
}
return res

}
//leetcode submit region end(Prohibit modification and deletion)
