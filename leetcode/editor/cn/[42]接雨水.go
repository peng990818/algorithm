//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
// 
//
// 示例 2： 
//
// 
//输入：height = [4,2,0,3,2,5]
//输出：9
// 
//
// 
//
// 提示： 
//
// 
// n == height.length 
// 1 <= n <= 2 * 10⁴ 
// 0 <= height[i] <= 10⁵ 
// 
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5384 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 比较暴力
// func trap(height []int) int {
//     if len(height) <= 1 {
//         return 0
//     }
//     count := 0
//     for i:=1;i<len(height)-1;i++ {
//         j, k := i-1, i+1
//         maxL, maxR := 0, 0
//         for j>=0 {
//             if maxL < height[j] {
//                 maxL = height[j]
//             }
//             j--
//         }
//         for k<len(height) {
//             if maxR < height[k] {
//                 maxR = height[k]
//             }
//             k++
//         }
//         min := maxL
//         if min > maxR {
//             min = maxR
//         }
//         if min > height[i] {
//             count += min-height[i]
//         }
//     }
//     return count
// }

// 使用动态规划进行优化
// func trap(height []int) int {
//     n := len(height)
//     if n == 0 {
//         return 0
//     }
//     leftMax := make([]int, n)
//     leftMax[0] = height[0]
//     for i:=1;i<n;i++ {
//         leftMax[i] = max(leftMax[i-1], height[i])
//     }
//
//     rightMax := make([]int, n)
//     rightMax[n-1] = height[n-1]
//     for j:=n-2;j>=0;j-- {
//         rightMax[j] = max(rightMax[j+1], height[j])
//     }
//
//     count := 0
//     for i, h := range height {
//         count += min(leftMax[i], rightMax[i]) - h
//     }
//     return count
// }

// 使用双指针来优化空间
// func trap(height []int) int {
//     left, right := 0, len(height)-1
//     leftMax, rightMax := 0, 0
//     res := 0
//     for left < right {
//         leftMax = max(leftMax, height[left])
//         rightMax = max(rightMax, height[right])
//         if height[left] < height[right] {
//             res += leftMax-height[left]
//             left++
//         } else {
//             res += rightMax-height[right]
//             right--
//         }
//     }
//     return res
// }

// 使用单调栈的方法
func trap(height []int) (ans int) {
stack := []int{}
for i, h := range height {
for len(stack) > 0 && h > height[stack[len(stack)-1]] {
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]
if len(stack) == 0 {
break
}
left := stack[len(stack)-1]
curWidth := i - left - 1
curHeight := min(height[left], h) - height[top]
ans += curWidth * curHeight
}
stack = append(stack, i)
}
return
}

func min(a, b int) int {
if a < b {
return a
}
return b
}

//leetcode submit region end(Prohibit modification and deletion)
