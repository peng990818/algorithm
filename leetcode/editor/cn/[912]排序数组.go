//给你一个整数数组 nums，请你将该数组升序排列。 
//
// 你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
// 
//
// 示例 2： 
//
// 
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 5 * 10⁴ 
// -5 * 10⁴ <= nums[i] <= 5 * 10⁴ 
// 
//
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 1058 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 堆排
// func sortArray(nums []int) []int {
//     if len(nums) <= 1 {
//         return nums
//     }
//     buildHeap(nums)
//     for size:=len(nums)-1;size > 0;size-- {
//         nums[size], nums[0] = nums[0], nums[size]
//         heapify(nums, 0, size)
//     }
//     return nums
// }
//
// func buildHeap(arr []int) {
//     for i:=len(arr)/2-1;i>=0;i-- {
//         heapify(arr, i, len(arr))
//     }
// }
//
// func heapify(arr []int, index, size int) {
//     left := 2*index+1
//     for left < size {
//         largest := index
//         if arr[largest] < arr[left] {
//             largest = left
//         }
//         if left+1 < size && arr[largest] < arr[left+1] {
//             largest = left+1
//         }
//         if largest == index {
//             break
//         }
//         arr[largest], arr[index] = arr[index], arr[largest]
//         index = largest
//         left = 2*index + 1
//     }
// }

// 快排
// func sortArray(nums []int) []int {
//     quickSort(nums, 0, len(nums)-1)
//     return nums
// }
//
// func quickSort(nums []int, l, r int) {
//     if l < r {
//         p := partition(nums, l, r)
//         quickSort(nums, l, p[0]-1)
//         quickSort(nums, p[1]+1, r)
//     }
// }
//
// func partition(nums []int, l, r int) []int {
//     less, more := l-1, r
//     for l < more {
//         if nums[l] < nums[r] {
//             less++
//             nums[l], nums[less] = nums[less], nums[l]
//             l++
//         } else if nums[l] > nums[r] {
//             more--
//             nums[l], nums[more] = nums[more], nums[l]
//         } else {
//             l++
//         }
//     }
//     nums[more], nums[r] = nums[r], nums[more]
//     return []int{less+1, more}
// }

// 归并
func sortArray(nums []int) []int {
    mergeSort(nums, 0, len(nums)-1)
    return nums
}

func mergeSort(nums []int, l, r int) {
    if l < r {
        mid := l + (r-l)/2
        mergeSort(nums, l, mid)
        mergeSort(nums, mid+1, r)
        merge(nums, l, mid, r)
    }
}

func merge(nums []int, l, mid, r int) {
    help := make([]int, r-l+1)
    i, j, k := l, mid+1, 0
    for i <= mid && j <= r {
        if nums[i] < nums[j] {
            help[k] = nums[i]
            i++
        } else {
            help[k] = nums[j]
            j++
        }
        k++
    }
    for i<=mid {
        help[k] = nums[i]
        i++
        k++
    }
    for j<=r {
        help[k] = nums[j]
        j++
        k++
    }
    for i:=0;i<len(help);i++ {
        nums[l+i] = help[i]
    }
}
//leetcode submit region end(Prohibit modification and deletion)
