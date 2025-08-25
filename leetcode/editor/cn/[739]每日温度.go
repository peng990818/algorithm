//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现
//在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。 
//
// 
//
// 示例 1: 
//
// 
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
// 
//
// 示例 2: 
//
// 
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
// 
//
// 示例 3: 
//
// 
//输入: temperatures = [30,60,90]
//输出: [1,1,0] 
//
// 
//
// 提示： 
//
// 
// 1 <= temperatures.length <= 10⁵ 
// 30 <= temperatures[i] <= 100 
// 
//
// Related Topics 栈 数组 单调栈 👍 1849 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

// func dailyTemperatures(temperatures []int) []int {
//     n := len(temperatures)
//     res := make([]int, n)
//     stack := []int{}
//     for i:=0;i<n;i++ {
//         t := temperatures[i]
//         for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
//             prevIndex := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             res[prevIndex] = i-prevIndex
//         }
//         stack = append(stack, i)
//     }
//     return res
// }

// func dailyTemperatures(temperatures []int) []int {
// if len(temperatures) == 0 {
// return nil
// }
// stack := make([]int, 0)
// stack = append(stack, 0)
// res := make([]int, len(temperatures))
// for i := 1; i < len(temperatures); i++ {
// if len(stack) == 0 || temperatures[i] < temperatures[stack[len(stack)-1]] {
// stack = append(stack, i)
// continue
// }
// for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
// top := stack[len(stack)-1]
// res[top] = i - top
// stack = stack[:len(stack)-1]
// }
// stack = append(stack, i)
// }
// return res
// }

func dailyTemperatures(temperatures []int) []int {
if len(temperatures) == 0 {
return nil
}
stack := make([]int, 0)
res := make([]int, len(temperatures))
for i := 0; i < len(temperatures); i++ {
for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
top := stack[len(stack)-1]
res[top] = i - top
stack = stack[:len(stack)-1]
}
stack = append(stack, i)
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
