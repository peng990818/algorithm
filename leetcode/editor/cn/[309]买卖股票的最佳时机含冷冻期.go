//给定一个整数数组
// prices，其中第 prices[i] 表示第 i 天的股票价格 。 
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）: 
//
// 
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。 
// 
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 
//
// 示例 1: 
//
// 
//输入: prices = [1,2,3,0,2]
//输出: 3 
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出] 
//
// 示例 2: 
//
// 
//输入: prices = [1]
//输出: 0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= prices.length <= 5000 
// 0 <= prices[i] <= 1000 
// 
//
// Related Topics 数组 动态规划 👍 1780 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 分析每一个节点可能所处的状态
// 把每个状态的值计算出来进行比较
// func maxProfit(prices []int) int {
//     if len(prices) == 0 {
//         return 0
//     }
//     n := len(prices)
//     dp := make([][3]int, len(prices))
//     dp[0][0] = -prices[0]
//     for i:=1;i<n;i++ {
//         dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
//         dp[i][1] = dp[i-1][0] + prices[i]
//         dp[i][2] = max(dp[i-1][1], dp[i-1][2])
//     }
//     return max(dp[n-1][1], dp[n-1][2])
// }

func maxProfit(prices []int) int {
if len(prices) == 0 {
return 0
}
n := len(prices)
f0, f1, f2 := -prices[0], 0, 0
for i := 1; i < n; i++ {
newf0 := max(f0, f2 - prices[i])
newf1 := f0 + prices[i]
newf2 := max(f1, f2)
f0, f1, f2 = newf0, newf1, newf2
}
return max(f1, f2)
}

func max(x, y int) int {
if x > y {
return x
}
return y
}

//leetcode submit region end(Prohibit modification and deletion)
