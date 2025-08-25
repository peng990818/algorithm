//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。 
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。 
//
// 你可以认为每种硬币的数量是无限的。 
//
// 
//
// 示例 1： 
//
// 
//输入：coins = [1, 2, 5], amount = 11
//输出：3 
//解释：11 = 5 + 5 + 1 
//
// 示例 2： 
//
// 
//输入：coins = [2], amount = 3
//输出：-1 
//
// 示例 3： 
//
// 
//输入：coins = [1], amount = 0
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= coins.length <= 12 
// 1 <= coins[i] <= 2³¹ - 1 
// 0 <= amount <= 10⁴ 
// 
//
// Related Topics 广度优先搜索 数组 动态规划 👍 2926 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func coinChange(coins []int, amount int) int {
//     // todo 待深入
//     if len(coins) == 0 || amount < 0 {
//         return 0
//     }
//     dp := make([]int, amount+1)
//     for i:=1;i<len(dp);i++ {
//         dp[i] = amount+1
//     }
//     for i:=0;i<=amount;i++ {
//         for j := 0;j<len(coins);j++ {
//             if coins[j] <= i {
//                 dp[i] = min(dp[i], dp[i-coins[j]]+1)
//             }
//         }
//     }
//     if dp[amount] > amount {
//         return -1
//     }
//     return dp[amount]
// }

// func coinChange(coins []int, amount int) int {
//     if len(coins) == 0 || amount <= 0 {
//         return 0
//     }
//     dp, max := make([]int, amount+1), amount+1
//     dp[0] = 0
//     for i:=1;i<max;i++ {
//         dp[i] = max
//     }
//     for i:=1;i<=amount;i++ {
//         for j:=0;j<len(coins);j++ {
//             if coins[j] <= i {
//                 dp[i] = min(dp[i], dp[i-coins[j]]+1)
//             }
//         }
//     }
//     if dp[amount] > amount {
//         return -1
//     }
//     return dp[amount]
// }

// func coinChange(coins []int, amount int) int {
// n := len(coins)
// dp := make([][]int, n+1)
// for i:=0;i<=n;i++ {
// dp[i] = make([]int, amount+1)
// dp[i][0] = 0
// }
// for j:=1;j<=amount;j++ {
// dp[0][j] = math.MaxInt32
// }
//
// for i:=1;i<=n;i++ {
// for j:=1;j<=amount;j++ {
// dp[i][j] = dp[i-1][j]
// if j >= coins[i-1] {
// dp[i][j] = min(dp[i][j-coins[i-1]]+1, dp[i][j])
// }
// }
// }
// if dp[n][amount] == math.MaxInt32 {
// return -1
// }
// return dp[n][amount]
// }

func coinChange(coins []int, amount int) int {
dp := make([]int, amount+1)
// 初始化，金额为零，最少需要0个货币
dp[0] = 0
for i:=1;i<=amount;i++ {
dp[i] = math.MaxInt32
}

for i:=0;i<len(coins);i++ {
for j:=coins[i];j<=amount;j++ {
dp[j] = min(dp[j-coins[i]]+1, dp[j])
}
}
if dp[amount] == math.MaxInt32 {
return -1
}
return dp[amount]
}
//leetcode submit region end(Prohibit modification and deletion)
