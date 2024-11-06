//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。 
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。 
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。 
//
// 
//
// 示例 1: 
//
// 
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
// 
//
// 示例 2: 
//
// 
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4 
//
// 
//
// 提示： 
//
// 
// 1 <= k <= nums.length <= 10⁵ 
// -10⁴ <= nums[i] <= 10⁴ 
// 
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2585 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
    buildHeap(nums)
    for size := len(nums)-1;size > 0 && k > 1;size-- {
        nums[size], nums[0] = nums[0], nums[size]
        heapify(nums, 0, size)
        k--
    }
    return nums[0]
}

func buildHeap(nums []int) {
    n := len(nums)
    for i:=n/2-1;i>=0;i-- {
        heapify(nums, i, n)
    }
}

func heapify(nums []int, index, size int) {
    left := 2*index+1
    for left < size {
        largest := index
        if nums[largest] < nums[left] {
            largest = left
        }
        if left+1 < size && nums[largest] < nums[left+1] {
            largest = left+1
        }
        if largest == index {
            break
        }
        nums[largest], nums[index] = nums[index], nums[largest]
        index = largest
        left = 2*index+1
    }
}
//leetcode submit region end(Prohibit modification and deletion)
