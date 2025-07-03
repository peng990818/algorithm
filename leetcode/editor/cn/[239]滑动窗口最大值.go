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

// type myStruct struct {
//     index int
//     val int
// }
//
// func maxSlidingWindow(nums []int, k int) []int {
//     if len(nums) < k {
//         return nil
//     }
//     h := make([]myStruct, k)
//     for i:=0;i<k;i++ {
//         h[i].index = i
//         h[i].val = nums[i]
//     }
//     buildHeap(h)
//     res := make([]int, 0, len(nums)-k+1)
//     res = append(res, h[0].val)
//     for i:=k;i<len(nums);i++ {
//         push(&h, i, nums[i])
//         for h[0].index <= i-k {
//             pop(&h)
//         }
//         res = append(res, h[0].val)
//     }
//     return res
// }
//
// func push(arr *[]myStruct, index, val int) {
//     *arr = append(*arr, myStruct{index: index, val: val})
//     heapifyUp(*arr, len(*arr)-1)
// }
//
// func heapifyUp(arr []myStruct, index int) {
//     p := (index-1)/2
//     for index > 0 && arr[index].val > arr[p].val {
//         arr[index], arr[p] = arr[p], arr[index]
//         index = p
//         p = (index-1)/2
//     }
// }
//
// func pop(arr *[]myStruct) myStruct {
//     top := (*arr)[0]
//     (*arr)[0], (*arr)[len(*arr)-1] = (*arr)[len(*arr)-1], (*arr)[0]
//     *arr = (*arr)[:len(*arr)-1]
//     heapify(*arr, 0, len(*arr))
//     return top
// }
//
// func buildHeap(arr []myStruct) {
//     for i:=len(arr)/2-1;i>=0;i-- {
//         heapify(arr, i, len(arr))
//     }
// }
//
// func heapify(arr []myStruct, index, size int) {
//     left := 2*index + 1
//     for left < size {
//         largest := index
//         if arr[largest].val < arr[left].val {
//             largest = left
//         }
//         if left+1<size && arr[largest].val < arr[left+1].val {
//             largest = left+1
//         }
//         if largest == index {
//             break
//         }
//         arr[largest], arr[index] = arr[index], arr[largest]
//         index = largest
//         left = 2*index+1
//     }
// }

func maxSlidingWindow(nums []int, k int) []int {
// 存储最大值, 单调队列
queue := make([]int, 0)
for i := 0; i < k && i < len(nums); i++ {
for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
queue = queue[:len(queue)-1]
}
queue = append(queue, nums[i])

}
res := []int{}
if len(queue) > 0 {
res = append(res, queue[0])
}
for i := k; i < len(nums); i++ {
if len(queue) > 0 && nums[i-k] == queue[0] {
queue = queue[1:]
}
for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
queue = queue[:len(queue)-1]
}
queue = append(queue, nums[i])
res = append(res, queue[0])
}
return res
}

//leetcode submit region end(Prohibit modification and deletion)
