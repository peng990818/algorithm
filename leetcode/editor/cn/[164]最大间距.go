//给定一个无序的数组 nums，返回 数组在排序之后，相邻元素之间最大的差值 。如果数组元素个数小于 2，则返回 0 。 
//
// 您必须编写一个在「线性时间」内运行并使用「线性额外空间」的算法。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [3,6,9,1]
//输出: 3
//解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。 
//
// 示例 2: 
//
// 
//输入: nums = [10]
//输出: 0
//解释: 数组元素个数小于 2，因此返回 0。 
//
// 
//
// 提示: 
//
// 
// 1 <= nums.length <= 10⁵ 
// 0 <= nums[i] <= 10⁹ 
// 
//
// Related Topics 数组 桶排序 基数排序 排序 👍 637 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func maximumGap(nums []int) int {
//     if len(nums) <= 1 {
//         return 0
//     }
//     sort.Ints(nums)
//     res := 0
//     for i:=1;i<len(nums);i++ {
//         res = max(res, nums[i]-nums[i-1])
//     }
//     return res
// }

// 基数排序
func maximumGap(nums []int) (ans int) {
    if len(nums) <= 1 {
        return 0
    }
    nums = radixSort(nums)
    for i:=1;i<len(nums);i++ {
        ans = max(ans, nums[i]-nums[i-1])
    }
    return
}

func radixSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    max := findMax(nums)
    exp := 1
    for max/exp > 0 {
        countSort(nums, exp)
        exp*=10
    }
    return nums
}

func findMax(nums []int) int {
    max := nums[0]
    for i:=1;i<len(nums);i++ {
        if max < nums[i] {
            max = nums[i]
        }
    }
    return max
}

func countSort(nums []int, exp int) {
    n := len(nums)
    output := make([]int, n)
    cnt := make([]int, 10)
    for i:=0;i<len(nums);i++ {
        index := (nums[i]/exp) % 10
        cnt[index]++
    }
    for i:=1;i<10;i++ {
        cnt[i] += cnt[i-1]
    }

    for i:=n-1;i>=0;i-- {
        index := (nums[i]/exp) % 10
        output[cnt[index]-1] = nums[i]
        cnt[index]--
    }
    copy(nums, output)
}

//leetcode submit region end(Prohibit modification and deletion)
