//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。 
//
// 返回 滑动窗口中的最大值 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], k = 1
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// -10⁴ <= nums[i] <= 10⁴ 
// 1 <= k <= nums.length 
// 
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2932 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func maxSlidingWindow(nums []int, k int) []int {
//     if len(nums) < k {
//         return nil
//     }
//     l, r := 0, k-1
//     res := make([]int, 0)
//     for r < len(nums) {
//         res = append(res, maxNum(nums[l:r+1]))
//         l++
//         r++
//     }
//     return res
// }
//
// func maxNum(nums []int) int {
//     max := math.MinInt32
//     for _, v := range nums {
//         if max < v {
//             max = v
//         }
//     }
//     return max
// }

// var a []int
// type hp struct{ sort.IntSlice }
// func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
// func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
// func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
//
// func maxSlidingWindow(nums []int, k int) []int {
// a = nums
// q := &hp{make([]int, k)}
// for i := 0; i < k; i++ {
// q.IntSlice[i] = i
// }
// heap.Init(q)
//
// n := len(nums)
// ans := make([]int, 1, n-k+1)
// ans[0] = nums[q.IntSlice[0]]
// for i := k; i < n; i++ {
// heap.Push(q, i)
// for q.IntSlice[0] <= i-k {
// heap.Pop(q)
// }
// ans = append(ans, nums[q.IntSlice[0]])
// }
// return ans
// }





//leetcode submit region end(Prohibit modification and deletion)
