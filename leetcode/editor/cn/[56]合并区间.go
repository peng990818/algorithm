//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。 
//
// 
//
// 示例 1： 
//
// 
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 
//
// 示例 2： 
//
// 
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。 
//
// 
//
// 提示： 
//
// 
// 1 <= intervals.length <= 10⁴ 
// intervals[i].length == 2 
// 0 <= starti <= endi <= 10⁴ 
// 
//
// Related Topics 数组 排序 👍 2439 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    for i:=0;i<len(intervals)-1;i++ {
        for j:=0;j<len(intervals)-1-i;j++ {
            if intervals[j][0] > intervals[j+1][0] {
                intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
            }
        }
    }
    res := make([][]int, 0)
    tmp, tmps := intervals[0], [][]int{}
    for i:=1;i<len(intervals);i++ {
        tmps = process(tmp, intervals[i])
        if len(tmps) == 2 {
            res = append(res, tmp)
            // 没产生新合并
            tmp = tmps[1]
        } else {
            // 产生新合并
            tmp = tmps[0]
        }
    }
    res = append(res, tmp)
    return res
}

func process(a []int, b []int) [][]int {
    if a[0] <= b[0] && b[0] <= a[1] {
        if a[1] <= b[1] {
            return [][]int{{a[0], b[1]}}
        }
        return [][]int{a}
    }
    if b[0] <= a[0] && a[0] <= b[1] {
        if b[1] <= a[1] {
            return [][]int{{b[0], a[1]}}
        }
        return [][]int{b}
    }
    return [][]int{a, b}
}
//leetcode submit region end(Prohibit modification and deletion)
