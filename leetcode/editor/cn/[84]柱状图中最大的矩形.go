//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。 
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
// 
//
// 示例 2： 
//
// 
//
// 
//输入： heights = [2,4]
//输出： 4 
//
// 
//
// 提示： 
//
// 
// 1 <= heights.length <=10⁵ 
// 0 <= heights[i] <= 10⁴ 
// 
//
// Related Topics 栈 数组 单调栈 👍 2825 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
n := len(heights)
left, right := make([]int, n), make([]int, n)
for i := 0; i < n; i++ {
right[i] = n
}
mono_stack := []int{}
for i := 0; i < n; i++ {
for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
right[mono_stack[len(mono_stack)-1]] = i
mono_stack = mono_stack[:len(mono_stack)-1]
}
if len(mono_stack) == 0 {
left[i] = -1
} else {
left[i] = mono_stack[len(mono_stack)-1]
}
mono_stack = append(mono_stack, i)
}
ans := 0
for i := 0; i < n; i++ {
ans = max(ans, (right[i] - left[i] - 1) * heights[i])
}
return ans
}

func max(x, y int) int {
if x > y {
return x
}
return y
}

//leetcode submit region end(Prohibit modification and deletion)
