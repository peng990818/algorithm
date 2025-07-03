//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
// 
//
// 示例 2: 
//
// 
//输入: nums = [1], k = 1
//输出: [1] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// k 的取值范围是 [1, 数组中不相同的元素的个数] 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的 
// 
//
// 
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。 
//
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 1915 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func topKFrequent(nums []int, k int) []int {
// occurrences := map[int]int{}
// for _, num := range nums {
// occurrences[num]++
// }
// h := &IHeap{}
// heap.Init(h)
// for key, value := range occurrences {
// heap.Push(h, [2]int{key, value})
// if h.Len() > k {
// heap.Pop(h)
// }
// }
// ret := make([]int, k)
// for i := 0; i < k; i++ {
// ret[k - i - 1] = heap.Pop(h).([2]int)[0]
// }
// return ret
// }
//
// type IHeap [][2]int
//
// func (h IHeap) Len() int           { return len(h) }
// func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
// func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
// func (h *IHeap) Push(x interface{}) {
// *h = append(*h, x.([2]int))
// }
//
// func (h *IHeap) Pop() interface{} {
// old := *h
// n := len(old)
// x := old[n-1]
// *h = old[0 : n-1]
// return x
// }

// func topKFrequent(nums []int, k int) []int {
//     mp := make(map[int]int)
//     for _, v := range nums {
//         if _, ok := mp[v]; ok {
//             mp[v]++
//             continue
//         }
//         mp[v] = 1
//     }
//     res := make([]int, 0, k)
//     for i:=0;i<k;i++ {
//         top, times := 0, 0
//         for key, v := range mp {
//             if v > times {
//                 times = v
//                 top = key
//             }
//         }
//         delete(mp, top)
//         res = append(res, top)
//     }
//     return res
// }

// func topKFrequent(nums []int, k int) []int {
//     if len(nums) == 0 {
//         return nil
//     }
//     mp := make(map[int]int, len(nums))
//     for _, v := range nums {
//         mp[v]++
//     }
//     arr := make([][]int, 0, len(mp))
//     for key, val := range mp {
//         arr = append(arr, []int{key, val})
//     }
//     buildHeap(arr)
//     res := make([]int, 0, k)
//     for i:=0;i<k;i++ {
//         res = append(res, pop(&arr))
//     }
//     return res
// }
//
// func heapify(arr [][]int, index int, size int) {
//     left := 2*index+1
//     for left < size {
//         largest := index
//         if arr[largest][1] < arr[left][1] {
//             largest = left
//         }
//         if left + 1 < size && arr[largest][1] < arr[left+1][1] {
//             largest = left+1
//         }
//         if largest == index {
//             break
//         }
//         index = largest
//         left = 2*index+1
//     }
// }
//
// func buildHeap(arr [][]int) {
//     for i:=len(arr)/2-1;i>=0;i-- {
//         heapify(arr, i, len(arr))
//     }
// }
//
// func pop(arr *[][]int) (res int) {
// if len(*arr) == 0 {
// return
// }
// res = (*arr)[0][0] // 堆顶元素
// (*arr)[0], (*arr)[len(*arr)-1] = (*arr)[len(*arr)-1], (*arr)[0]
// *arr = (*arr)[:len(*arr)-1] // 移除堆顶
// heapify(*arr, 0, len(*arr)) // 重新调整堆
// return
// }

// func topKFrequent(nums []int, k int) []int {
// if len(nums) == 0 {
// return nil
// }
//
// // 统计频率
// mp := make(map[int]int, len(nums))
// for _, v := range nums {
// mp[v]++
// }
//
// // 转换为二维数组 [值, 频率]
// arr := make([][]int, 0, len(mp))
// for key, val := range mp {
// arr = append(arr, []int{key, val})
// }
//
// // 构建堆
// buildHeap(arr)
//
// // 取前 k 个
// res := make([]int, 0, k)
// for i := 0; i < k; i++ {
// res = append(res, pop(&arr))
// }
//
// return res
// }
//
// func heapify(arr [][]int, index int, size int) {
// left := 2*index + 1
// for left < size {
// largest := index
// if arr[largest][1] < arr[left][1] {
// largest = left
// }
// if left+1 < size && arr[largest][1] < arr[left+1][1] {
// largest = left + 1
// }
// if largest == index {
// break
// }
// // 交换
// arr[index], arr[largest] = arr[largest], arr[index]
// index = largest
// left = 2*index + 1
// }
// }
//
// func buildHeap(arr [][]int) {
// for i := len(arr)/2 - 1; i >= 0; i-- {
// heapify(arr, i, len(arr))
// }
// }
//
// func pop(arr *[][]int) (res int) {
// if len(*arr) == 0 {
// return
// }
// res = (*arr)[0][0] // 堆顶元素
// (*arr)[0], (*arr)[len(*arr)-1] = (*arr)[len(*arr)-1], (*arr)[0]
// *arr = (*arr)[:len(*arr)-1] // 移除堆顶
// heapify(*arr, 0, len(*arr)) // 重新调整堆
// return
// }

func topKFrequent(nums []int, k int) []int {
mp := map[int]int{}
for _, n := range nums {
mp[n]++
}

arr := make([][]int, 0, len(mp))
for k, v := range mp {
arr = append(arr, []int{k, v})
}

buildHeap(arr)
res := []int{}
for i := 0; i < k; i++ {
res = append(res, popMax(&arr))
}
return res
}

func heapify(arr [][]int, index int, size int) {
left := 2*index + 1
for left < size {
largest := index
if arr[largest][1] < arr[left][1] {
largest = left
}
if left+1 < size && arr[largest][1] < arr[left+1][1] {
largest = left + 1
}
if largest == index {
return
}
arr[largest], arr[index] = arr[index], arr[largest]
index = largest
left = 2*index + 1
}
}

func buildHeap(arr [][]int) {
for i := len(arr)/2 - 1; i >= 0; i-- {
heapify(arr, i, len(arr))
}
}

func popMax(arr *[][]int) int {
res := (*arr)[0][0]
(*arr)[0], (*arr)[len(*arr)-1] = (*arr)[len(*arr)-1], (*arr)[0]
*arr = (*arr)[:len(*arr)-1]
heapify(*arr, 0, len(*arr))
return res
}



//leetcode submit region end(Prohibit modification and deletion)
