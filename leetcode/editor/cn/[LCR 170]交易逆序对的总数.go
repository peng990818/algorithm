//在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的
//「交易逆序对」总数。 
//
// 
//
// 示例 1: 
//
// 
//输入：record = [9, 7, 5, 4, 6]
//输出：8
//解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
// 
//
// 
//
// 限制： 
//
// 0 <= record.length <= 50000 
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1119 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(record []int) int {
    return process(record, 0, len(record)-1)
}

func process(nums []int, l, r int) int {
    if l < r {
        mid := l + (r-l)/2
        return process(nums, l, mid) + process(nums, mid+1, r) +
            merge(nums, l, mid, r)
    }
    return 0
}

func merge(nums []int, l, mid, r int) int {
    help := make([]int, r-l+1)
    i, j, k := l, mid+1, 0
    res := 0
    for i<=mid && j<=r {
        if nums[i] <= nums[j] {
            help[k] = nums[i]
            i++
        } else {
            help[k] = nums[j]
            // 重点
            res += mid-i+1
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
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
